package utils

import "strconv"

func ConvertUint8(value string) uint8 {
	res, _ := strconv.Atoi(value)
	return uint8(res)
}

func ConvertInt32(value string) int32 {
	res, _ := strconv.Atoi(value)
	return int32(res)
}

func ConvertInt(value string) int {
	res, _ := strconv.Atoi(value)
	return res
}

func ConvertBool(value string) bool {
	res, _ := strconv.ParseBool(value)
	return res
}
