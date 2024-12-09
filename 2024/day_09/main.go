package main

import (
	_024 "github.com/endertunc/advent-of-code/2024"
	"slices"
)

// 6201130364722
func part1(diskMap []int) int {
	dotIndex := slices.Index(diskMap, -1)
	for i := len(diskMap) - 1; i >= 0; i-- {
		n := diskMap[i]
		if n == -1 {
			continue
		}
		dotIndex = findNextEmptySpace(diskMap, dotIndex)
		if dotIndex >= i {
			break
		}
		diskMap[dotIndex] = diskMap[i]
		diskMap[i] = -1
	}
	return findSum(diskMap)
}

// 6221662795602
func part2(diskMap []int) int {
	currentNumber, currenNumberIndex, currentNumberLength := findCurrentNumberSequence(diskMap, len(diskMap)-1)
	for {
		for {
			// find the first empty space that can fit the current number sequence
			spaceIndex, _ := findSpaceWithLSequenceLength(diskMap, 0, currentNumberLength)
			// if there is no space or the space is after the current number sequence, we are done
			if spaceIndex == -1 || spaceIndex > currenNumberIndex {
				break
			} else {
				// fill the empty space with the current number sequence
				for i := 0; i < currentNumberLength; i++ {
					diskMap[spaceIndex+i] = currentNumber
				}
				// mark the current number sequence location as empty space
				for i := currentNumberLength - 1; i >= 0; i-- {
					diskMap[currenNumberIndex-i] = -1
				}
				break
			}
		}
		// find the next current number sequence
		currentNumber, currenNumberIndex, currentNumberLength = findCurrentNumberSequence(diskMap, currenNumberIndex-currentNumberLength)
		// if the current number sequence is after the first empty space, we are done
		if slices.Index(diskMap, -1) >= currenNumberIndex { // not sure if there is a better way to track "first empty space"
			break
		}
	}

	return findSum(diskMap)
}

func findSum(diskMap []int) int {
	sum := 0
	for i, value := range diskMap {
		if value == -1 {
			continue
		}
		sum += value * i
	}
	return sum
}

func findNextEmptySpace(s []int, startIndex int) int {
	for i := startIndex; i < len(s); i++ {
		if s[i] == -1 {
			return i
		}
	}
	return -1
}

func findCurrentNumberSequence(s []int, startIndex int) (int, int, int) {
	currentNumber := 0
	currentNumberIndex := 0
	// skip empty spaces first and find the first non-empty number
	for i := startIndex; i >= 0; i-- {
		if s[i] != -1 {
			currentNumber = s[i]
			currentNumberIndex = i
			break
		}
	}

	// find the index and length of the current number sequence
	for i := currentNumberIndex; i >= 0; i-- {
		if s[i] != currentNumber {
			return currentNumber, currentNumberIndex, currentNumberIndex - i
		}
	}
	return -1, -1, -1
}

func findSpaceWithLSequenceLength(s []int, startIndex int, minLength int) (int, int) {
	spaceIndex := findNextEmptySpace(s, startIndex) // skip non-space characters
	for i := spaceIndex; i < len(s); i++ {
		if s[i] != -1 {
			if i-spaceIndex >= minLength {
				return spaceIndex, i - spaceIndex
			} else {
				// we might have found an empty space, but it is not long enough,
				// so we try to find a longer one starting from where we left (startIndex = i)
				return findSpaceWithLSequenceLength(s, i, minLength)
			}
		}
	}
	return -1, -1
}

func main() {
	filename := "inputs/day_09/input.txt"
	input := _024.ReadInput(filename)

	println(part1(parseInput(input)))
	println(part2(parseInput(input)))
}

func parseInput(input string) []int {
	diskMap := make([]int, 0)
	for i := 0; i < len(input); i++ {
		n := _024.MustParseInt(string(input[i]))
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				diskMap = append(diskMap, i/2)
			}
		} else {
			for j := 0; j < n; j++ {
				diskMap = append(diskMap, -1) // empty spaces represented as -1
			}
		}
	}
	return diskMap
}
