package handler

import (
	"log"
	"net/http"

	"github-proxy/static"
)

// IndexHandler 处理首页请求
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] 处理首页请求: %s", r.URL.String())

	// 检查是否有查询参数q
	q := r.URL.Query().Get("q")
	if q != "" {
		log.Printf("[INFO] 首页请求带有查询参数q=%s，重定向", q)
		http.Redirect(w, r, "/"+q, http.StatusFound)
		return
	}

	// 设置内容类型并返回HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(static.IndexHTML))
}

// FaviconHandler 处理图标请求
//func FaviconHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "image/vnd.microsoft.icon")
//	w.Write(static.IconData)
//}
