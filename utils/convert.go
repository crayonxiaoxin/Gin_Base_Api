package utils

import (
	"strconv"
)

func Str2Int(str string, def ...int) int {
	numInt64, err := strconv.ParseInt(str, 0, 0)
	numInt := int(numInt64)
	if err != nil {
		if len(def) > 0 {
			numInt = def[0]
		} else {
			numInt = 0
		}
	}
	return numInt
}
