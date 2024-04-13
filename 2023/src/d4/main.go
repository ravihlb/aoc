package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	winningNumbers []int
	gotNumbers     []int
}

func main() {
	const filepath string = "./d4.input"

    answer := solve(filepath)
    fmt.Println("Total points sum:", answer)
}

func solve(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	cards := parseCards(file)
	pointsPerCard := calcPointsPerCard(cards)
    // fmt.Println("Got points:", pointsPerCard)

    var totalPointsSum int
    for _, points := range pointsPerCard {
        totalPointsSum += points
    }

    return totalPointsSum
}

func calcPointsPerCard(cards []Card) []int {
    var pointsPerCard []int

    for _, card := range cards {
        var points int
        var matches int

        for _, winningNumber := range card.winningNumbers {
            for _, gotNumber := range card.gotNumbers {
                if gotNumber == winningNumber {
                    matches += 1
                }
            }
        }

        for i := 1; i <= matches; i++ {
            if points == 0 {
                points = 1
                continue
            }

            points *= 2
        }

        pointsPerCard = append(pointsPerCard, points)
    }

	return pointsPerCard
}

func parseCards(file *os.File) []Card {
	var cards []Card
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		cardNumPattern := regexp.MustCompile("Card [0-9]: ")
		lineUnnamed := cardNumPattern.ReplaceAllString(line, "")
		lineUnnamedSplit := strings.Split(lineUnnamed, " | ")

		winningNumbersString := strings.Split(lineUnnamedSplit[0], " ")
		winningNumbers := sliceAtoi(winningNumbersString)

		gotNumbersString := strings.Split(lineUnnamedSplit[1], " ")
		gotNumbers := sliceAtoi(gotNumbersString)

		cards = append(cards, Card{
			winningNumbers: winningNumbers,
			gotNumbers:     gotNumbers},
		)
	}

	return cards
}

func sliceAtoi(slice []string) []int {
	var out []int

	for _, val := range slice {
		num, err := strconv.Atoi(val)

		if err != nil {
			continue
		}

		out = append(out, num)
	}

	return out
}
