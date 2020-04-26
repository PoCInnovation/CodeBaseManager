package codebase

func FoundAllArgs(found []bool) bool {
	for _, value := range found {
		if !value {
			return false
		}
	}
	return true
}
