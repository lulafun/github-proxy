package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Config 应用配置
type Config struct {
	Host      string
	Port      int
	Debug     bool
	Timeout   time.Duration
	JsDelivr  bool
	SizeLimit int64
	ChunkSize int
	WhiteList [][]string
	BlackList [][]string
	PassList  [][]string
}

// GetConfig 从环境变量加载配置
func GetConfig() *Config {
	return &Config{
		Host:      getEnv("GH_PROXY_HOST", "0.0.0.0"),
		Port:      getEnvAsInt("GH_PROXY_PORT", 8080),
		Debug:     getEnvAsBool("GH_PROXY_DEBUG", false),
		Timeout:   time.Duration(getEnvAsInt("GH_PROXY_TIMEOUT", 3600)) * time.Second,
		JsDelivr:  getEnvAsBool("GH_PROXY_JSDELIVR", false),
		SizeLimit: getEnvAsInt64("GH_PROXY_SIZE_LIMIT", 1024*1024*1024*999),
		ChunkSize: getEnvAsInt("GH_PROXY_CHUNK_SIZE", 1024*10),
		WhiteList: parseList(getEnv("GH_PROXY_WHITE_LIST", "")),
		BlackList: parseList(getEnv("GH_PROXY_BLACK_LIST", "")),
		PassList:  parseList(getEnv("GH_PROXY_PASS_LIST", "")),
	}
}

// 获取环境变量
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 获取整数环境变量
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// 获取int64环境变量
func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}

// 获取布尔环境变量
func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		// 处理 "0"/"1" 的情况
		if valueStr == "0" {
			return false
		} else if valueStr == "1" {
			return true
		}
		return defaultValue
	}
	return value
}

// 解析规则列表
func parseList(list string) [][]string {
	var result [][]string
	if list == "" {
		return result
	}

	lines := strings.Split(list, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, "/")
		for i, part := range parts {
			parts[i] = strings.TrimSpace(part)
		}
		if len(parts) > 0 && parts[0] != "" {
			result = append(result, parts)
		}
	}
	return result
}
