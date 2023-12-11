package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	dayFlag := flag.Int("d", 1, "select which day (1-25)")
	partFlag := flag.Int("p", 1, "select which part (1-2)")
	flag.Parse()

	if *dayFlag < 1 || *dayFlag > 25 {
		panic(fmt.Errorf("day: %v is out of bounds", *dayFlag))
	}

	if *partFlag < 1 || *partFlag > 2 {
		panic(fmt.Errorf("part: %v is out of bounds", *partFlag))
	}

	switch *dayFlag {
	case 1:
		if *partFlag == 1 {
			day1Part1()
		} else {
			day1Part2()
		}
	case 2:
		if *partFlag == 1 {
			day2Part1()
		} else {
			day2Part2()
		}
	}
}

func day1Part1() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += day1.Part1(scanner.Bytes())
	}

	fmt.Println(sum)
}

func day1Part2() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += day1.Part2(scanner.Bytes())
	}

	fmt.Println(sum)
}

func day2Part1() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += day2.Part1(scanner.Bytes())
	}

	fmt.Println(sum)
}

func day2Part2() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += day2.Part2(scanner.Bytes())
	}

	fmt.Println(sum)
}
