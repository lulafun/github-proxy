# GitHub Proxy

<div align="center">
  <p><strong>Fast and reliable proxy for GitHub resources</strong></p>
</div>

[English](#overview) | [中文](#概述)

---

## Overview

GitHub Proxy is a lightweight Go service that proxies GitHub resources to enhance accessibility and download speeds. It's particularly useful in regions where GitHub access is restricted or slow.

### Features

- **Fast proxying** of GitHub repositories, raw files, releases, and archives
- **CDN acceleration** with optional jsDelivr integration
- **Access control** with customizable white lists and black lists
- **Multi-platform support** with pre-built binaries for various operating systems
- **Simple configuration** using environment variables
- **Streaming transfer** for efficient handling of large files
- **Google-inspired UI** with a clean and intuitive interface

### Screenshots

<div align="center">
  <img src="https://i.miji.bid/2025/04/06/f53924722a6a4b52e9d86e8463886400.png" alt="GitHub Proxy Screenshot" width="600"/>
</div>

## Installation

### Pre-built Binaries

Download the latest binary for your platform from the [Releases](https://github.com/yourusername/github-proxy/releases) page.

### Docker

```sh
docker run -d -p 8080:8080 -e GH_PROXY_HOST=0.0.0.0 lulafun/github-proxy
```

### Building from Source

```sh
git clone https://github.com/lulafun/github-proxy.git
cd github-proxy
go build
```

## Configuration

GitHub Proxy is configured via environment variables:

| Variable | Description                                            | Default                 |
|----------|--------------------------------------------------------|-------------------------|
| `GH_PROXY_HOST` | Host address to listen on                              | `127.0.0.1`             |
| `GH_PROXY_PORT` | Port to listen on                                      | `8080`                  |
| `GH_PROXY_DEBUG` | Enable debug logging                                   | `false`                 |
| `GH_PROXY_TIMEOUT` | Timeout for proxy(second)                              | `3600 second`           |
| `GH_PROXY_JSDELIVR` | Enable jsDelivr CDN for blob/raw files                 | `false`                 |
| `GH_PROXY_SIZE_LIMIT` | Maximum file size (bytes)                              | `1073741824000` (999GB) |
| `GH_PROXY_CHUNK_SIZE` | Stream chunk size (bytes)                              | `10240` (10KB)          |
| `GH_PROXY_WHITE_LIST` | List of allowed repositories (newline separated)       | ``                      |
| `GH_PROXY_BLACK_LIST` | List of blocked repositories (newline separated)       | ``                      |
| `GH_PROXY_PASS_LIST` | List of repos to directly redirect (newline separated) | ``                      |

### Access Control Format

Each line in the white/black/pass lists can be in one of these formats:
```
username
username/repo
*/repo
```

## Usage

1. Start the proxy server:
   ```sh
   ./github-proxy
   ```

2. Access GitHub resources through the proxy:
   ```
   http://localhost:8080/github.com/golang/go/archive/master.zip
   http://localhost:8080/github.com/golang/go/blob/master/README.md
   http://localhost:8080/raw.githubusercontent.com/golang/go/master/README.md
   ```

3. Or use the web interface at `http://localhost:8080`

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## References
[gh-proxy](https://github.com/hunshcn/gh-proxy)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

# GitHub 代理服务

## 概述

GitHub 代理是一个轻量级的 Go 服务，用于代理 GitHub 资源以提高访问性和下载速度。它在 GitHub 访问受限或速度较慢的地区特别有用。

### 特性

- **快速代理** GitHub 仓库、原始文件、发布包和存档
- **CDN 加速**，可选择性地集成 jsDelivr
- **访问控制**，使用可自定义的白名单和黑名单
- **多平台支持**，为各种操作系统提供预构建的二进制文件
- **简单配置**，使用环境变量
- **流式传输**，高效处理大文件
- **Google 风格的 UI**，干净直观的界面

### 截图

<div align="center">
  <img src="https://i.miji.bid/2025/04/06/f53924722a6a4b52e9d86e8463886400.png" alt="GitHub 代理截图" width="600"/>
</div>

## 安装

### 预构建二进制文件

从[发布页面](https://github.com/lulafun/github-proxy/releases)下载适合您平台的最新二进制文件。

### Docker

```sh
docker run -d -p 8080:8080 -e GH_PROXY_HOST=0.0.0.0 lulafun/github-proxy
```

### 从源码构建

```sh
git clone https://github.com/lulafun/github-proxy.git
cd github-proxy
go build
```

## 配置

GitHub 代理通过环境变量进行配置：

| 变量 | 描述                           | 默认值                     |
|----------|------------------------------|-------------------------|
| `GH_PROXY_HOST` | 监听的主机地址                      | `127.0.0.1`             |
| `GH_PROXY_PORT` | 监听的端口                        | `8080`                  |
| `GH_PROXY_DEBUG` | 启用调试日志                       | `false`                 |
| `GH_PROXY_TIMEOUT` | 代理超时时间                       | `3600 秒`                |
| `GH_PROXY_JSDELIVR` | 为 blob/raw 文件启用 jsDelivr CDN | `false`                 |
| `GH_PROXY_SIZE_LIMIT` | 最大文件大小（字节）                   | `1073741824000` (999GB) |
| `GH_PROXY_CHUNK_SIZE` | 流传输块大小（字节）                   | `10240` (10KB)          |
| `GH_PROXY_WHITE_LIST` | 允许的仓库列表（换行分隔）                | ``                      |
| `GH_PROXY_BLACK_LIST` | 禁止的仓库列表（换行分隔）                | ``                      |
| `GH_PROXY_PASS_LIST` | 直接重定向的仓库列表（换行分隔）             | ``                      |

### 访问控制格式

白名单/黑名单/通过列表中的每一行可以是以下格式之一：
```
用户名
用户名/仓库
*/仓库
```

## 使用方法

1. 启动代理服务器：
   ```sh
   ./github-proxy
   ```

2. 通过代理访问 GitHub 资源：
   ```
   http://localhost:8080/github.com/golang/go/archive/master.zip
   http://localhost:8080/github.com/golang/go/blob/master/README.md
   http://localhost:8080/raw.githubusercontent.com/golang/go/master/README.md
   ```

3. 或者使用 `http://localhost:8080` 的网页界面

## 贡献

欢迎贡献！请随时提交 Pull Request。

## 参考
[gh-proxy](https://github.com/hunshcn/gh-proxy)

## 许可证

本项目采用 MIT 许可证 - 详情请查看 [LICENSE](LICENSE) 文件。