package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// regex to match mul(ddd,ddd) 1 - 3 digits

func solvePartOne(scanner *bufio.Scanner) {
	var allMatches []string
	for scanner.Scan() {
		line := scanner.Text()
		// check for matches and append to slice
		matches := getMatches(line)
		allMatches = append(allMatches, matches...)
	}

	// extract values from matches and multiply
	fmt.Println(getSum(allMatches))

}

func getSum(matches []string) int {
	sum := 0
	pattern := `\d{1,3},\d{1,3}`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

	for i := range matches {
		values := re.FindAllString(matches[i], -1)
		for _, v := range values {
				// split v by comma
				split := strings.Split(v, ",")
				if len(split) == 2 {
					v1, err := strconv.Atoi(split[0])
					if err != nil {
						fmt.Println("Error parsing string", err)
					}
					v2, err := strconv.Atoi(split[1])
					if err != nil {
						fmt.Println("Error parsing string", err)
					}
					sum += v1 * v2
				}
		}
	}

	return sum
}

func getMatches(line string) []string {
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

	matches := re.FindAllString(line, -1)
	return matches
}

func main() {
	file, err := os.Open("./day3input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	solvePartOne(scanner)



	// scan corrupted memory for uncorruped mul instructions (mul(num, num))

	// get the sum of all valid mul instructions

	// return sum
}
