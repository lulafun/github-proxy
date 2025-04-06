package util

// MatchList 检查是否匹配规则列表
func MatchList(m []string, list [][]string) bool {
	if len(m) < 2 {
		return false
	}

	for _, item := range list {
		if len(item) == 1 {
			// 只有用户名的规则
			if m[0] == item[0] {
				return true
			}
		} else if len(item) >= 2 {
			// 用户名/仓库名的规则或通配符规则
			if item[0] == "*" && m[1] == item[1] {
				return true
			} else if m[0] == item[0] && m[1] == item[1] {
				return true
			}
		}
	}

	return false
}
