package util

// InArray 查找某值是否在数组中
func InArray(v string, m []string) bool {
	for _, value := range m {
		if value == v {
			return true
		}
	}
	return false
}
