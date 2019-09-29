package sample

// InStatusList checks if the arg is in the list
func InStatusList(x string) bool {
	ls := []string{"drafted", "published", "deleted"}
	for _, s := range ls {
		if x == s {
			return true
		}
	}
	return false
}
