package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of code day 2!")
	file, err := os.Open("./day2Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numSafeRows := 0
	index := 0
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		numIntList, err := getRowOfNums(nums)
		if err != nil {
			log.Fatal(err)
		}
		var safeOrNot bool
		inc := false
		dec := false
		for index := range numIntList {
			if index == 0 {
				continue
			} else {
				distance := numIntList[index] - numIntList[index-1]
				if index == 1 {
					if distance < 0 {
						inc = true
					} else {
						dec = true
					}
				}

				if int(math.Abs(float64(distance))) > 3 || int(math.Abs(float64(distance))) < 1 {
					safeOrNot = false
					break
				} else {
					if (inc && distance > 0) || (dec && distance < 0) {
						safeOrNot = false
						break
					} else {
						safeOrNot = true
					}
				}
			}
		}
		if safeOrNot {
			fmt.Printf("Row %d is safe\n", index)
			numSafeRows++
		} else {
			fmt.Printf("Row %d is not safe\n", index)
		}
		index++
	}
	fmt.Printf("Number of safe rows: %d\n", numSafeRows)
}

func getRowOfNums(numStrings []string) ([]int, error) {
	var numRowList []int

	for _, element := range numStrings {
		num, err := strconv.Atoi(element)
		if err != nil {
			return nil, err
		}
		numRowList = append(numRowList, num)
	}

	return numRowList, nil
}
