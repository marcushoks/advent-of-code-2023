package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := [][]byte{
		[]byte("1abc2"),
		[]byte("pqr3stu8vwx"),
		[]byte("a1b2c3d4e5f"),
		[]byte("treb7uchet"),
	}

	var sum int
	for _, bytes := range input {
		sum += Part1(bytes)
	}

	expected := 142
	if sum != expected {
		t.Errorf("Expected %v, got %v", expected, sum)
	}
}
