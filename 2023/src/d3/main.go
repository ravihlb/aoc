package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Vector struct {
	X, Y int
}

type Gear struct {
	symbol         string
	adjacentNumber int
	position       Vector
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	const inputFilePath = "./d3.input"

	var (
		finalPartsSum int
		gearRatios    []int
		gearRatiosSum int
	)

	file, err := os.Open(inputFilePath)
	check(err)

	partNumbers, gears := getPartNumbers(file)

	for _, partNumber := range partNumbers {
		finalPartsSum += partNumber
	}

	gearRatios = getGearRatios(gears)
	for _, ratio := range gearRatios {
		gearRatiosSum += ratio
	}

	fmt.Println("Final part numbers sum:", finalPartsSum)
	fmt.Println("Final gear ratios sum:", gearRatiosSum)
}

func getCharVectorsFromCurrIdx(lineNr, lineNrMax, charIdx, charIdxMax int) (out []Vector) {
	directions := map[string]Vector{
		"Up":        {charIdx, lineNr - 1},
		"Down":      {charIdx, lineNr + 1},
		"UpLeft":    {charIdx - 1, lineNr - 1},
		"Left":      {charIdx - 1, lineNr},
		"DownLeft":  {charIdx - 1, lineNr + 1},
		"UpRight":   {charIdx + 1, lineNr - 1},
		"Right":     {charIdx + 1, lineNr},
		"DownRight": {charIdx + 1, lineNr + 1},
	}

	for _, vector := range directions {
		if vector.X > -1 && vector.X < charIdxMax && vector.Y < lineNrMax && vector.Y > -1 {
			out = append(out, vector)
		}
	}

	return out
}

func isCharInt(char string) bool {
	_, err := strconv.Atoi(char)

	if err != nil {
		return false
	}

	return true
}

func loadFileContent(file *os.File) []string {
	var content []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content
}

func getGearRatios(gears []Gear) []int {
	gearRatios := []int{}
	hits := 0

	for needleIdx, needle := range gears {
		hits = 0
		var foundGear Gear

		for gearIdx, gear := range gears {
			if hits > 1 {
				break
			}

			if needleIdx == gearIdx || gear.adjacentNumber == needle.adjacentNumber {
				continue
			}

			if gear.position == needle.position {
				foundGear = gear
				hits++
			}
		}

		if hits == 1 {
			fmt.Println("gearRatio :=", needle.adjacentNumber, "*", foundGear.adjacentNumber)

			gearRatio := needle.adjacentNumber * foundGear.adjacentNumber
			fmt.Println("gearRatio", gearRatio)
			gearRatios = append(gearRatios, gearRatio)
		}

        gears = gears[1:]
	}

	return gearRatios
}

func getPartNumbers(file *os.File) (partNumbers []int, gears []Gear) {
	partNumbers = []int{}
	gears = []Gear{}

	content := loadFileContent(file)
	lineNrMax := len(content)

	for lineNr, lineContent := range content {
		var vectors []Vector

		charIdxMax := len(lineContent)
		isPartNumber := false
		currDigits := ""

	lineCharsLoop:
		for charIdx, charRune := range lineContent {
			var partSymbol Gear

			currChar := string(charRune)
			nextChar := ""

			if !isCharInt(currChar) {
				continue lineCharsLoop
			}

			currDigits += currChar
			vectors = append(vectors, getCharVectorsFromCurrIdx(
				lineNr,
				lineNrMax,
				charIdx,
				charIdxMax,
			)...)

			// avoid going out of line bounds
			if !((charIdx + 1) >= len(lineContent)) {
				nextChar = string(lineContent[charIdx+1])
			}

			// if nextChar is not empty, accumulate it
			if isCharInt(nextChar) {
				continue lineCharsLoop
			}

			currPartNumber, err := strconv.Atoi(currDigits)
			check(err)

			// iterate over surrounding chars to check for symbols
		vectorsLoop:
			for _, vector := range vectors {
				target := string(content[vector.Y][vector.X])

				if target == "." || isCharInt(target) {
					continue vectorsLoop
				} else {

					isPartNumber = true
					if target == "*" {
						partSymbol = Gear{target, currPartNumber, vector}
					}

					break vectorsLoop
				}
			}

			// reset vectors
			vectors = []Vector{}

			// by now we've checked every available nextChar and any adjacent symbols
			// if it is not a part number, trash it and move on
			if !isPartNumber {
				currDigits = ""
				continue lineCharsLoop
			}

			// appending our stuff
			partNumbers = append(partNumbers, currPartNumber)

			if partSymbol != (Gear{}) {
				gears = append(gears, partSymbol)
			}

			// reset our stuff for next iteration
			currDigits = ""
			isPartNumber = false
		}
	}

	return partNumbers, gears
}
