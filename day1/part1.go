package day1

import (
	"strconv"
)

func Part1(bytes []byte) int {
	var strDigits []string

	for _, b := range bytes {
		if b < 48 || b > 57 {
			continue
		}

		strDigits = append(strDigits, string(b))
	}

	i1, err1 := strconv.Atoi(strDigits[0])
	if err1 != nil {
		panic(err1)
	}

	i2, err2 := strconv.Atoi(strDigits[len(strDigits)-1])
	if err2 != nil {
		panic(err2)
	}

	return i1*10 + i2
}
