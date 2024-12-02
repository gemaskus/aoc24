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
		fmt.Printf("first num: %d ", firstNum)
		fmt.Printf("second num: %d\n", secondNum)
		firstNumList = append(firstNumList, firstNum)
		secondNumList = append(secondNumList, secondNum)
	}
	sort.Slice(firstNumList, func(i, j int) bool {
		return firstNumList[i] < firstNumList[j]
	})
	sort.Slice(secondNumList, func(i, j int) bool {
		return secondNumList[i] < secondNumList[j]
	})

	var distanceEachElement []int
	distanceSum := 0
	for index := range firstNumList {
		distance := int(math.Abs(float64(firstNumList[index]) - float64(secondNumList[index])))
		distanceEachElement = append(distanceEachElement, distance)
		distanceSum += distance
	}

	fmt.Println("Day 1, part 1:")
	fmt.Printf("Total Distance: %d\n", distanceSum)

	var simularityCount []int
	totalSimularityCount := 0
	for _, firstListElement := range firstNumList {
		simCount := 0
		for _, secondListElement := range secondNumList {
			if firstListElement == secondListElement {
				simCount++
			}
		}
		simCount *= firstListElement
		fmt.Printf("simularity score: %d\n", simCount)
		simularityCount = append(simularityCount, simCount)
		totalSimularityCount += simCount
	}
	fmt.Println("Day 1 part 2:")
	fmt.Printf("Total simularity count: %d\n", totalSimularityCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
