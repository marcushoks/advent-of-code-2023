package day2

import "testing"

func TestDay2(t *testing.T) {
	input := [][]byte{
		[]byte("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"),
		[]byte("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"),
		[]byte("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"),
		[]byte("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"),
		[]byte("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"),
	}

	sum := 0
	for _, bytes := range input {
		sum += Part2(bytes)
	}

	expected := 2286
	if sum != expected {
		t.Errorf("Expected: %v, got: %v", expected, sum)
	}
}
