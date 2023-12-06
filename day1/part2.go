package day1

import (
	"fmt"
	"regexp"
	"strconv"
)

func Part2(bytes []byte) int {
	r := regexp.MustCompile("1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine")
	matches := r.FindAllIndex(bytes, -1)

	firstStrDigit := string(bytes[matches[0][0]:matches[0][1]])
	lastStrDigit := string(bytes[matches[len(matches)-1][0]:matches[len(matches)-1][1]])

	firstDigit, err1 := convertStrToInt(firstStrDigit)
	if err1 != nil {
		panic(err1)
	}

	lastDigit, err2 := convertStrToInt(lastStrDigit)
	if err2 != nil {
		panic(err2)
	}

	return firstDigit*10 + lastDigit
}

func convertStrToInt(s string) (int, error) {
	var res int
	var err error

	if len(s) == 1 {
		return strconv.Atoi(s)
	}

	switch s {
	case "one":
		res = 1
	case "two":
		res = 2
	case "three":
		res = 3
	case "four":
		res = 4
	case "five":
		res = 5
	case "six":
		res = 6
	case "seven":
		res = 7
	case "eight":
		res = 8
	case "nine":
		res = 9
	default:
		err = fmt.Errorf("couldn't convert %v into integer", s)
	}

	return res, err
}
