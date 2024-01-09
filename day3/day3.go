package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./input")
	re := regexp.MustCompile(`^[^\w.]+$`)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input = [][]string{}

	var num string

	var sum int

	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
	}

	added := false

	for rowIndex, row := range input {
		for colIndex, col := range row {
			if col >= "0" && col <= "9" {
				num += col
				if colIndex+1 < len(row) && row[colIndex+1] == "." || colIndex == len(row)-1 {
					checkIndex := colIndex - len(num)

					//check if symbol on the same row but before it
					if checkIndex >= 0 && checkIndex < len(row) {
						if re.MatchString(row[checkIndex]) {
							val, err := strconv.Atoi(num)
							if err != nil {
								fmt.Println("Error converting string to int:", err)
								continue
							}
							fmt.Println("val:", val, rowIndex, colIndex)
							sum += val

							added = true
						}
					}

					//check if symbol  above it
					for i := 1; i <= len(num); i++ {
						pointerIndex := colIndex - len(num) + i
						if pointerIndex >= 0 && checkIndex < len(row) && rowIndex > 0 && num != "" {
							if re.MatchString(input[rowIndex-1][pointerIndex]) {
								val, err := strconv.Atoi(num)
								if err != nil {
									fmt.Println("Error converting string to int:", err)
									continue
								}

								fmt.Println("val:", val, rowIndex, colIndex)
								sum += val
								added = true

							}
						}

					}

					//check if symbol  below  it
					for i := 1; i <= len(num); i++ {
						pointerIndex := colIndex - len(num) + i
						if pointerIndex >= 0 && checkIndex < len(row) && rowIndex < len(input)-1 && num != "" {
							if re.MatchString(input[rowIndex+1][pointerIndex]) {
								val, err := strconv.Atoi(num)
								if err != nil {
									fmt.Println("Error converting string to int:", err)
									continue
								}

								fmt.Println("val:", val, rowIndex, colIndex)
								sum += val

								added = true
							}
						}

					}

					//check if symbol  top left of  it
					topLeftPointer := colIndex - len(num)
					if topLeftPointer >= 0 && rowIndex > 0 && num != "" {
						if re.MatchString(input[rowIndex-1][topLeftPointer]) {
							val, err := strconv.Atoi(num)
							if err != nil {
								fmt.Println("Error converting string to int:", err)
								continue
							}

							fmt.Println("val:", val, rowIndex, colIndex)
							sum += val

							added = true
						}

					}

					//check if symbol  top right of  it
					topRightPointer := colIndex + 1
					if topRightPointer < len(row) && rowIndex > 0 && num != "" {
						if re.MatchString(input[rowIndex-1][topRightPointer]) {
							val, err := strconv.Atoi(num)
							if err != nil {
								fmt.Println("Error converting string to int:", err)
								continue
							}

							fmt.Println("val:", val, rowIndex, colIndex)
							sum += val

							added = true
						}

					}

					//check if symbol bottom left of  it
					bottomLeft := colIndex - len(num)
					if bottomLeft >= 0 && rowIndex < len(input)-1 && num != "" {
						if re.MatchString(input[rowIndex+1][bottomLeft]) {
							val, err := strconv.Atoi(num)
							if err != nil {
								fmt.Println("Error converting string to int:", err)
								continue
							}

							fmt.Println("val:", val, rowIndex, colIndex)
							sum += val

							added = true
						}

					}

					//check if symbol bottom right of  it
					bottomRight := colIndex + 1
					if bottomRight < len(row) && rowIndex < len(input)-1 && num != "" {
						if re.MatchString(input[rowIndex+1][bottomRight]) {
							val, err := strconv.Atoi(num)
							if err != nil {
								fmt.Println("Error converting string to int:", err)
								continue
							}

							fmt.Println("val:", val, rowIndex, colIndex)
							sum += val

							added = true
						}
					}

				}

				if colIndex+1 < len(row) && re.MatchString(row[colIndex+1]) && num != "" {
					val, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println("Error converting string to int:", err)
						continue
					}

					fmt.Println("val:", val, rowIndex, colIndex)
					sum += val
					added = true

				}

				if added {
					num = ""
                    added = false
				}
			}
			if col == "." {
				num = ""
			}
		}

	}

	fmt.Printf("sum: %v\n", sum)

}
