package converter

import "strconv"

func ToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
