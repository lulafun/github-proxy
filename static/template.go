package static

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed assets
var staticFiles embed.FS

// 提供已处理的HTML内容和静态文件服务
var (
	IndexHTML  string
	FileServer http.Handler
)

func init() {
	// 读取HTML文件
	htmlBytes, err := fs.ReadFile(staticFiles, "assets/index.html")
	if err != nil {
		log.Fatalf("无法加载HTML: %v", err)
	}

	// 使用静态HTML内容
	IndexHTML = string(htmlBytes)

	// 为嵌入式文件系统创建HTTP文件服务器
	// 需要去掉"assets"前缀，以便正确提供文件
	stripped, err := fs.Sub(staticFiles, "assets")
	if err != nil {
		log.Fatalf("无法创建子文件系统: %v", err)
	}

	FileServer = http.FileServer(http.FS(stripped))
}
