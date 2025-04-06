package util

import (
	"regexp"
	"strings"
)

// URL正则表达式
var (
	Exp1 = regexp.MustCompile(`^(?:https?://)?github\.com/(?P<author>.+?)/(?P<repo>.+?)/(?:releases|archive)/.*$`)
	Exp2 = regexp.MustCompile(`^(?:https?://)?github\.com/(?P<author>.+?)/(?P<repo>.+?)/(?:blob|raw)/.*$`)
	Exp3 = regexp.MustCompile(`^(?:https?://)?github\.com/(?P<author>.+?)/(?P<repo>.+?)/(?:info|git-).*$`)
	Exp4 = regexp.MustCompile(`^(?:https?://)?raw\.(?:githubusercontent|github)\.com/(?P<author>.+?)/(?P<repo>.+?)/.+?/.+$`)
	Exp5 = regexp.MustCompile(`^(?:https?://)?gist\.(?:githubusercontent|github)\.com/(?P<author>.+?)/.+?/.+$`)
)

// CheckURL 检查URL是否为GitHub资源URL，并提取作者和仓库名
func CheckURL(u string) (bool, []string) {
	for _, exp := range []*regexp.Regexp{Exp1, Exp2, Exp3, Exp4, Exp5} {
		matches := exp.FindStringSubmatch(u)
		if matches != nil && len(matches) >= 3 {
			return true, matches[1:3]
		}
	}
	return false, nil
}

// FixURL 确保URL格式正确
func FixURL(url string) string {
	// 添加协议前缀
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		// 如果包含 s:/ 但不是 s:// 则修复双斜杠
		if strings.HasPrefix(url, "https:/") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url[7:]
		} else if strings.HasPrefix(url, "http:/") && !strings.HasPrefix(url, "http://") {
			url = "http://" + url[6:]
		} else {
			// 默认使用HTTPS
			url = "https://" + url
		}
	}

	return url
}

// ConvertBlobToRaw 将blob URL转换为raw URL
func ConvertBlobToRaw(url string) string {
	if Exp2.MatchString(url) {
		return strings.Replace(url, "/blob/", "/raw/", 1)
	}
	return url
}

// ConvertToJsdelivr 将URL转换为jsDelivr CDN URL
func ConvertToJsdelivr(url string) string {
	if Exp2.MatchString(url) {
		// 转换GitHub blob URL
		return strings.Replace(
			strings.Replace(url, "/blob/", "@", 1),
			"github.com", "cdn.jsdelivr.net/gh", 1)
	} else if Exp4.MatchString(url) {
		// 转换raw.githubusercontent.com URL
		re := regexp.MustCompile(`(\.com/.*?/.+?)/(.+?/)`)
		newURL := re.ReplaceAllString(url, "${1}@${2}")

		// 替换域名
		if strings.Contains(newURL, "raw.githubusercontent.com") {
			newURL = strings.Replace(newURL, "raw.githubusercontent.com", "cdn.jsdelivr.net/gh", 1)
		} else if strings.Contains(newURL, "raw.github.com") {
			newURL = strings.Replace(newURL, "raw.github.com", "cdn.jsdelivr.net/gh", 1)
		}

		return newURL
	}

	return url
}
