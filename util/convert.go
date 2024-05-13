package util

import "strconv"

func UintToString(number uint) string {
	return strconv.FormatUint(uint64(number), 10)
}

func Float64ToString(number float64) string {
	return strconv.FormatFloat(number, 'f', 0, 64)
}
