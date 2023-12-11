package day2

func Part1(bytes []byte) int {
	MAX_RED := 12
	MAX_GREEN := 13
	MAX_BLUE := 14

	game := ParseGame(bytes)
	legalSamples := 0
	for _, s := range game.samples {
		if s.red > MAX_RED || s.green > MAX_GREEN || s.blue > MAX_BLUE {
			break
		}
		legalSamples++
	}

	if legalSamples == len(game.samples) {
		return game.id
	}
	return 0
}
