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

							num = ""
							continue
						}
					}

					//check if symbol  above it
					for i := 1; i <= len(num); i++ {
						pointerIndex := colIndex - len(num) + i
						if pointerIndex >= 0 && checkIndex < len(row) && rowIndex > 0 {
							if re.MatchString(input[rowIndex-1][pointerIndex]) {
								val, err := strconv.Atoi(num)
								if err != nil {
									fmt.Println("Error converting string to int:", err)
									continue
								}

								fmt.Println("val:", val, rowIndex, colIndex)
								sum += val

								num = ""
								continue
							}
						}
					}

					//check if symbol  below  it
					for i := 1; i <= len(num); i++ {
						pointerIndex := colIndex - len(num) + i
						if pointerIndex >= 0 && checkIndex < len(row) && rowIndex < len(input)-1 {
							if re.MatchString(input[rowIndex+1][pointerIndex]) {
								val, err := strconv.Atoi(num)
								if err != nil {
									fmt.Println("Error converting string to int:", err)
									continue
								}

								fmt.Println("val:", val, rowIndex, colIndex)
								sum += val

								num = ""
								continue
							}
						}
					}

					//check if symbol  top left of  it
					topLeftPointer := colIndex - len(num) + 1
					if topLeftPointer >= 0 && rowIndex > 0 {
						if re.MatchString(input[rowIndex-1][topLeftPointer]) {
							val, err := strconv.Atoi(num)
							if err != nil {
								fmt.Println("Error converting string to int:", err)
								continue
							}

							fmt.Println("val:", val, rowIndex, colIndex)
							sum += val

							num = ""
							continue
						}
					}

					//check if symbol  top right of  it
					topRightPointer := colIndex + 1
					if topRightPointer < len(row) && rowIndex > 0 {
						if re.MatchString(input[rowIndex-1][topRightPointer]) {
							val, err := strconv.Atoi(num)
							if err != nil {
								fmt.Println("Error converting string to int:", err)
								continue
							}

							fmt.Println("val:", val, rowIndex, colIndex)
							sum += val

							num = ""
							continue
						}
					}

					//check if symbol bottom left of  it
					bottomLeft := colIndex - len(num)
					if bottomLeft >= 0 && rowIndex < len(input)-1 {
						if re.MatchString(input[rowIndex+1][bottomLeft]) {
							val, err := strconv.Atoi(num)
							if err != nil {
								fmt.Println("Error converting string to int:", err)
								continue
							}

							fmt.Println("val:", val, rowIndex, colIndex)
							sum += val

							num = ""
							continue
						}
					}

					//check if symbol bottom right of  it
					bottomRight := colIndex + 1
					if bottomRight < len(row) && rowIndex < len(input)-1 {
						if re.MatchString(input[rowIndex+1][bottomRight]) {
							val, err := strconv.Atoi(num)
							if err != nil {
								fmt.Println("Error converting string to int:", err)
								continue
							}

							fmt.Println("val:", val, rowIndex, colIndex)
							sum += val

							num = ""
							continue
						}
					}

				}

				if colIndex+1 < len(row) && re.MatchString(row[colIndex+1]) {
					val, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println("Error converting string to int:", err)
						continue
					}

					fmt.Println("val:", val, rowIndex, colIndex)
					sum += val

					num = ""
					continue
				}

			}
			if col == "." {
				num = ""
			}
		}

	}

	fmt.Printf("sum: %v\n", sum)

}
