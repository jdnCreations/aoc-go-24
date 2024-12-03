package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func solvePartOne() {
	file, err := os.Open("./day1.txt")
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// create two lists of items, left and right
	left := []int{}
	right := []int{}
	diff := []int{}

	for scanner.Scan() {
		var l, r int
		values := strings.Fields(scanner.Text())
		fmt.Sscanf(values[0], "%d", &l)
		fmt.Sscanf(values[1], "%d", &r)
		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error: %s", err.Error())	
	}

	// // loop over lists
	for range left {
	 minLeft := getMin(&left)
	 minRight := getMin(&right)
	 dist := getDistance(minLeft, minRight)
	 diff = append(diff, dist)
	}
	
	sum := getSum(diff)	
	fmt.Printf("Total: %d\n", sum)
}

func getSum(slice []int) int {
	total := 0
	for i := range slice {
		total += slice[i]
	}
	return total
}

func getMin(slice *[]int) int {
	s := *slice
	var min = s[0]
	index := 0
	for i, val := range s {
		if val < min {
			min = val
			index = i
		}
	}

	*slice = append(s[:index], s[index+1:]...)

	return min
}

func getDistance(num1, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		return -diff
	}
	return diff
}

func main() {
	solvePartOne()
	solvePartTwo()
}