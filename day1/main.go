package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("AoC 2024 Day1!\n")
	file, err := os.Open("./day1Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var firstNumList []int
	var secondNumList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		firstNum, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		secondNum, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		firstNumList = append(firstNumList, firstNum)
		secondNumList = append(secondNumList, secondNum)
	}
	firstNumList = sortSlice(firstNumList)
	secondNumList = sortSlice(secondNumList)

	distanceSum := getTotalDistance(firstNumList, secondNumList)

	fmt.Println("Day 1, part 1:")
	fmt.Printf("Total Distance: %d\n", distanceSum)

	totalSimularityCount := getTotalSimularityScore(firstNumList, secondNumList)
	fmt.Println("Day 1 part 2:")
	fmt.Printf("Total simularity count: %d\n", totalSimularityCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
func sortSlice(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

func getTotalDistance(first []int, second []int) int {
	distanceSum := 0
	for index := range first {
		distance := int(math.Abs(float64(first[index]) - float64(second[index])))
		distanceSum += distance
	}
	return distanceSum
}

func getTotalSimularityScore(first []int, second []int) int {
	totalSimularityCount := 0
	for _, firstListElement := range first {
		simCount := 0
		for _, secondListElement := range second {
			if firstListElement == secondListElement {
				simCount++
			}
		}
		simCount *= firstListElement
		totalSimularityCount += simCount
	}
	return totalSimularityCount
}
