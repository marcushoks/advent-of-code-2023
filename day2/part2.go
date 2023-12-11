package day2

func Part2(bytes []byte) int {
	game := ParseGame(bytes)

	maxRed := -1
	maxGreen := -1
	maxBlue := -1

	for _, sample := range game.samples {
		if sample.red > maxRed {
			maxRed = sample.red
		}

		if sample.green > maxGreen {
			maxGreen = sample.green
		}

		if sample.blue > maxBlue {
			maxBlue = sample.blue
		}
	}

	return maxRed * maxGreen * maxBlue
}
