package handler

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github-proxy/config"
	"github-proxy/util"
)

// ProxyHandler 代理处理器
type ProxyHandler struct {
	Config *config.Config
}

// NewProxyHandler 创建代理处理器
func NewProxyHandler(cfg *config.Config) *ProxyHandler {
	return &ProxyHandler{Config: cfg}
}

// ServeHTTP 实现http.Handler接口
func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] 处理代理请求: %s %s", r.Method, r.URL.Path)

	path := r.URL.Path[1:] // 移除前导斜杠

	// 修复URL格式
	path = util.FixURL(path)

	// 检查URL格式
	valid, matches := util.CheckURL(path)
	if !valid {
		http.Error(w, "Invalid input.", http.StatusForbidden)
		return
	}

	// 检查白名单
	if len(h.Config.WhiteList) > 0 && !util.MatchList(matches, h.Config.WhiteList) {
		http.Error(w, "Forbidden by white list.", http.StatusForbidden)
		return
	}

	// 检查黑名单
	if util.MatchList(matches, h.Config.BlackList) {
		http.Error(w, "Forbidden by black list.", http.StatusForbidden)
		return
	}

	// 检查通过列表和jsDelivr转换 - 仅保留这部分判断，其他都由自动重定向处理
	passBy := util.MatchList(matches, h.Config.PassList)

	// 仅在启用jsDelivr或满足通过列表条件时才进行jsDelivr重定向
	if (h.Config.JsDelivr || passBy) && (util.Exp2.MatchString(path) || util.Exp4.MatchString(path)) {
		newPath := util.ConvertToJsdelivr(path)
		http.Redirect(w, r, newPath, http.StatusFound)
		return
	}

	// 对于blob URL转换为raw URL (GitHub特性)
	if util.Exp2.MatchString(path) {
		path = util.ConvertBlobToRaw(path)
	}

	// 处理通过列表的直接重定向
	if passBy {
		fullURL := path
		if r.URL.RawQuery != "" {
			fullURL += "?" + r.URL.RawQuery
		}

		fullURL = util.FixURL(fullURL)
		http.Redirect(w, r, fullURL, http.StatusFound)
		return
	}

	// 执行简化的代理请求
	h.proxyRequest(w, r, path)
}

// proxyRequest 执行简化的代理请求，自动处理重定向
func (h *ProxyHandler) proxyRequest(w http.ResponseWriter, r *http.Request, targetURL string) {
	// 构建完整URL
	fullURL := targetURL
	if r.URL.RawQuery != "" {
		fullURL += "?" + r.URL.RawQuery
	}

	// 修复URL格式
	fullURL = util.FixURL(fullURL)

	log.Printf("[INFO] 代理请求: %s %s", r.Method, fullURL)

	// 创建新请求
	req, err := http.NewRequest(r.Method, fullURL, r.Body)
	if err != nil {
		log.Printf("[ERROR] 创建请求失败: %v", err)
		http.Error(w, "Error creating request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//// 复制请求头
	//for name, values := range r.Header {
	//	if name != "Host" && name != "Connection" {
	//		for _, value := range values {
	//			req.Header.Add(name, value)
	//		}
	//	}
	//}

	// 创建自动处理重定向的客户端
	client := &http.Client{
		Timeout: h.Config.Timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 允许所有重定向，但记录日志
			if len(via) >= 10 {
				log.Printf("[WARN] 停止在10次重定向后")
				return http.ErrUseLastResponse
			}
			log.Printf("[INFO] 重定向到: %s", req.URL.String())
			return nil
		},
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] 请求失败: %v", err)
		http.Error(w, "Request failed: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	log.Printf("[INFO] 收到响应: %d %s", resp.StatusCode, resp.Status)

	// 检查文件大小限制
	if contentLength := resp.Header.Get("Content-Length"); contentLength != "" {
		if size, err := strconv.ParseInt(contentLength, 10, 64); err == nil {
			if size > h.Config.SizeLimit {
				log.Printf("[WARN] 文件大小超过限制: %d > %d", size, h.Config.SizeLimit)
				http.Redirect(w, r, fullURL, http.StatusFound)
				return
			}
			log.Printf("[INFO] 文件大小: %d bytes", size)
		}
	}

	// 复制响应头
	for name, values := range resp.Header {
		if name != "Connection" && name != "Transfer-Encoding" {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}
	}

	// 设置状态码
	w.WriteHeader(resp.StatusCode)

	// 如果是HEAD请求，不需要写入响应体
	if r.Method == http.MethodHead {
		return
	}

	// 流式写入响应体
	buf := make([]byte, h.Config.ChunkSize)
	written := int64(0)
	startTime := time.Now()

	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			written += int64(n)

			if _, writeErr := w.Write(buf[:n]); writeErr != nil {
				log.Printf("[ERROR] 写入响应失败: %v", writeErr)
				return
			}

			// 刷新缓冲区
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}

		if err != nil {
			if err != io.EOF {
				log.Printf("[ERROR] 读取响应失败: %v", err)
			}
			break
		}
	}

	duration := time.Since(startTime)
	bytesPerSec := float64(written) / duration.Seconds()

	log.Printf("[INFO] 代理完成: %d bytes, %.2f bytes/sec", written, bytesPerSec)
}
