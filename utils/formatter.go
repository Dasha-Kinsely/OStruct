package utils

import "strconv"

func FormatIntToString(integer int64) string {
	return strconv.FormatInt(integer, 10)
}

func FormatFloatToString(float float64) string {
	return strconv.FormatFloat(float, 64, 4, 64)
}