package common

import "strconv"

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
func ConvertToInt(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		return 0, err
	} else {
		return num, nil
	}
}

func ConvertToBool(str string) (bool, error) {
	num, err := strconv.ParseBool(str)

	if err != nil {
		return false, err
	} else {
		return num, nil
	}
}
