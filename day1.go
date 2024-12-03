package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func solvePartOne(scanner bufio.Scanner) {
	
	// create two lists of items, left and right
	left, right := createTwoLists(scanner)	
	diff := []int{}
	
	// // loop over lists
	for range left {
	 minLeft := getMin(&left)
	 minRight := getMin(&right)
	 dist := getDistance(minLeft, minRight)
	 diff = append(diff, dist)
	}
	
	sum := getSum(diff)	
	fmt.Printf("Part 1: %d\n", sum)
}

func solvePartTwo(scanner bufio.Scanner) {
	total := 0	
	left, right := createTwoLists(scanner)
	// // get first value of list 1.

	for i := range left {
		leftValue := left[i]
		count := countValues(leftValue, right)
		total += count * leftValue
	}

	fmt.Printf("Part 2: %d\n", total)
	// check how many times that value is in list 2
	// multiply first value by how many times its in second list, add to a new similarity list
}

func createTwoLists(scanner bufio.Scanner) ([]int,[]int) {
	left := []int{}
	right := []int{}

	fmt.Println(scanner.Text())

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

	return left, right
}

func countValues(value int, slice []int) int {
	total := 0
	for i := range slice {
		if slice[i] == value {
			total += 1
		}
	}
	return total
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
	file, err := os.Open("./day1.txt")
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	solvePartOne(*scanner)
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	solvePartTwo(*scanner)
}