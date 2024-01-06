package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []string
	var number string
	answer := 0

	for scanner.Scan() {
		for _, v := range strings.Split(scanner.Text(), "") {
			if _, err := strconv.Atoi(v); err == nil {
				number += v
			}

		}
		if len(number) > 2 {
			firstChar := string(number[0])
			lastChar := string(number[len(number)-1])
			number = firstChar + lastChar

		}

		if len(number) == 1 {
			number += number
		}
		numbers = append(numbers, number)
		number = ""
	}


	for _, v := range numbers {
		i, err := strconv.Atoi(v)
    
        answer += i
		if err != nil {
			panic(err)
		}

	}

    fmt.Println(answer)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
