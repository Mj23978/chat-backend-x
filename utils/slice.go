package utils

func Contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func Delete(slice []string, place int) []string {
	if len(slice) == 1 {
		slice[0] = ""
		return slice
	}
	slice[place] = slice[len(slice)-1]
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]
	return slice
}