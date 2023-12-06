package day2

import (
	"regexp"
)

func Part1(bytes []byte) int {
	MAX_RED := 12
	MAX_GREEN := 13
	MAX_BLUE := 14

	match := regexp.MustCompile("[0-9]+:").Find(bytes)
	gameId := match[:len(match)-1] // removes trailing ":"

}
