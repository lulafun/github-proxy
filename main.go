package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github-proxy/config"
	"github-proxy/handler"
	"github-proxy/static"
)

func init() {
	// 设置日志格式
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// 加载配置
	cfg := config.GetConfig()

	// 创建代理处理器
	proxyHandler := handler.NewProxyHandler(cfg)

	// 创建路由
	mux := http.NewServeMux()

	// 注册路由
	// 在main.go的路由注册部分
	//mux.HandleFunc("/favicon.ico", handler.FaviconHandler)

	// 添加对静态资源的处理
	mux.Handle("/static/", http.StripPrefix("/static/", static.FileServer))

	// 首页和代理处理保持不变
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handler.IndexHandler(w, r)
		} else {
			proxyHandler.ServeHTTP(w, r)
		}
	})

	// 创建服务器
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 启动服务器
	log.Printf("[INFO] 启动GitHub代理服务器，监听地址: %s:%d", cfg.Host, cfg.Port)
	log.Printf("[INFO] 配置: JsDelivr=%v, 调试=%v", cfg.JsDelivr, cfg.Debug)
	log.Fatal(server.ListenAndServe())
}
