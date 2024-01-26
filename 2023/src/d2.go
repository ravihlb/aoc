/*
Determine which games would have been possible if the bag had been loaded with
only 12 red cubes, 13 green cubes, and 14 blue cubes.
What is the sum of the IDs of those games?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type CubeGame struct {
	id            int
	maxRedShown   int
	maxBlueShown  int
	maxGreenShown int
}

func main() {
	const filePath = "./d2.input"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Could not read file", err)
		return
	}

	limits := CubeGame{
		maxRedShown:   12,
		maxGreenShown: 13,
		maxBlueShown:  14,
	}

	games, err := makeGamesList(file, limits)
	for i := 0; i < len(games); i++ {
		fmt.Println("Game id:", games[i].id)
		fmt.Println("max red shown:", games[i].maxRedShown)
		fmt.Println("max blue shown:", games[i].maxBlueShown)
		fmt.Println("max green shown:", games[i].maxGreenShown)
		fmt.Println()
	}

	var answer int
	fmt.Println(limits)

	for _, game := range games {

		fmt.Println()
		fmt.Println(
			"Game id:", game.id, "\n",
			"max red shown:", game.maxRedShown, "\n",
			"max blue shown:", game.maxBlueShown, "\n",
			"max green shown:", game.maxGreenShown,
		)
		fmt.Println()

		if game.maxRedShown <= limits.maxRedShown &&
			game.maxBlueShown <= limits.maxBlueShown &&
			game.maxGreenShown <= limits.maxGreenShown {

			fmt.Println("Game id", game.id, "is valid")
			answer += game.id
			continue
		}
		fmt.Println("Game id", game.id, "is invalid")
	}
	fmt.Println("Final sum:", answer)
}

func makeGamesList(file *os.File, limits CubeGame) ([]CubeGame, error) {
	var games []CubeGame
	scanner := bufio.NewScanner(file)

	// for each line
	for scanner.Scan() {
		var id int
		var err error

		var (
			maxRedShown   int
			maxGreenShown int
			maxBlueShown  int
		)

		line := scanner.Text()
		wordsStr := strings.Replace(line, ", ", " ", -1)
		wordsStr = strings.Replace(wordsStr, "; ", " ", -1)
		var words []string = strings.Fields(wordsStr)

		// for each word
		for i := 0; i < len(words); i++ {

			switch words[i] {
			case "Game":
				idStr := strings.Trim(words[i+1], ":")
				fmt.Println("idStr", idStr)

				id, err = strconv.Atoi(idStr)
				check(err)
				continue

			case "red":
				countStr := words[i-1]
				rCount, err := strconv.Atoi(countStr)
				check(err)

				if rCount > maxRedShown {
					maxRedShown = rCount
				}
				continue

			case "blue":
				countStr := words[i-1]
				bCount, err := strconv.Atoi(countStr)
				check(err)

				if bCount > maxBlueShown {
					maxBlueShown = bCount
				}
				continue

			case "green":
				countStr := words[i-1]
				gCount, err := strconv.Atoi(countStr)
				check(err)

				if gCount > maxGreenShown {
					maxGreenShown = gCount
				}
				continue

			}
		}

		game := CubeGame{
			id,
			maxRedShown,
			maxBlueShown,
			maxGreenShown,
		}

		fmt.Println("Just created game", game)
		games = append(games, game)
	}

	return games, nil
}
