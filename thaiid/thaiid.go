package thaiid

import (
	"strconv"
)

func ValidateThaiId(thaiid string) bool {
	if len(thaiid) != 13 {
		return false
	} else {
		length := len(thaiid)
		total := 0
		for i := 0; i < 12; i++ {
			num, _ := strconv.Atoi(string(thaiid[i]))
			total = total + (num * length)
			length -= 1
		}
		a := total % 11
		last, _ := strconv.Atoi(string(thaiid[len(thaiid)-1]))
		x := (11 - a) % 10
		if x == last {
			return true
		} else {
			return false
		}
	}
}
