package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var validIds []int
	sumIds := 0

	for scanner.Scan() {
		games := strings.Split(scanner.Text(), ":")
		ids := strings.Split(games[0], " ")
		idStr := ids[1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Error converting id to integer:", err)
		}
		rounds := games[1]
		turn := strings.Split(rounds, ",")

		validGames := []bool{true}
		for _, v := range turn {
			play := strings.Split(v, " ")
			fmt.Printf("play: %v\n", play)
			for i := 0; i < len(play); i++ {
				num, err := strconv.ParseInt(play[i], 10, 64)
				if err != nil {
					fmt.Printf("Non-numeric value: %v\n", play[i])
					continue
				}
				if (play[i+1] == "red" || play[i+1] == "red;") && num > 12 {
					validGames = append(validGames, false)
				}

				if (play[i+1] == "green" || play[i+1] == "green;") && num > 13 {
					validGames = append(validGames, false)
				}
				if (play[i+1] == "blue" || play[i+1] == "blue;") && num > 14 {
					validGames = append(validGames, false)
				}

			}

		}
		if !slices.Contains(validGames, false) {
			validIds = append(validIds, id)
		}
		validGames = nil

	}

	for i := 0; i < len(validIds); i++ {
		sumIds += validIds[i]
	}

	fmt.Print(sumIds)

}

func part2() {

	file, err := os.Open("./input2")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumOfProuct := 0

	for scanner.Scan() {
		games := strings.Split(scanner.Text(), ":")

		rounds := games[1]
		turn := strings.Split(rounds, ",")

		var maxRed, maxGreen, maxBlue int64

		for _, v := range turn {
			play := strings.Split(v, " ")
			fmt.Printf("play: %v\n", play)
			for i := 0; i < len(play); i++ {
				num, err := strconv.ParseInt(play[i], 10, 64)
				if err != nil {
					fmt.Printf("Non-numeric value: %v\n", play[i])
					continue
				}
				if (play[i+1] == "red" || play[i+1] == "red;") && num > maxRed {
					maxRed = num
				}

				if (play[i+1] == "green" || play[i+1] == "green;") && num > maxGreen {
					maxGreen = num
				}
				if (play[i+1] == "blue" || play[i+1] == "blue;") && num > maxBlue {
					maxBlue = num
				}

			}

		}
		product := maxRed * maxGreen * maxBlue
		sumOfProuct += int(product)
		product = 0

	}

	fmt.Println(sumOfProuct)

}

func main() {
	part2()
}
