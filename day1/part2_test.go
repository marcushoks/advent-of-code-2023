package day1

import (
	"testing"
)

func TestPart2(t *testing.T) {
	input := [][]byte{
		[]byte("two1nine"),
		[]byte("eightwothree"),
		[]byte("abcone2threexyz"),
		[]byte("xtwone3four"),
		[]byte("4nineeightseven2"),
		[]byte("zoneight234"),
		[]byte("7pqrstsixteen"),
	}

	var sum int
	for _, bytes := range input {
		sum += Part2(bytes)
	}

	expected := 281
	if sum != expected {
		t.Errorf("Expected %v, got %v", expected, sum)
	}
}

func TestPart2Overlap(t *testing.T) {
	input := []byte("twone")
	sum := Part2(input)
	expected := 21
	if sum != expected {
		t.Errorf("Expected %v, got %v", expected, sum)
	}
}
