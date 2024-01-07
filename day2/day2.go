package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//12red 13 green 14 blue
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var redCubes, greenCubes, blueCubes int
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

		for i, v := range rounds {
			if unicode.IsDigit(v) && i+2 < len(rounds) {
				digit := int(v - '0')

				if rounds[i+2] == 'r' {
					redCubes += digit
				}
				if rounds[i+2] == 'g' {
					greenCubes += digit
				}
				if rounds[i+2] == 'b' {
					blueCubes += digit
				}

			}
		}

		if redCubes <= 12 && greenCubes <= 13 && blueCubes <= 14 {
			sumIds += id
            fmt.Println(redCubes, greenCubes, blueCubes)
		}

		redCubes = 0
		greenCubes = 0
		blueCubes = 0

	}

	fmt.Print(sumIds)
}
