package helpers

import "strconv"

func Floattostr(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
