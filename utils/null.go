package utils

func StringNullCheck(value string) bool {
	if value != "" {
		return true
	}
	return false
}

func IntNullCheck(value int) bool {
	if value != 0 {
		return true
	}
	return false
}