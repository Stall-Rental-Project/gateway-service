package common

func ContainsString(items []string, ele string) bool {
	if items == nil {
		return false
	}

	for _, item := range items {
		if item == ele {
			return true
		}
	}

	return false
}
