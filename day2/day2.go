package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	utils "github.com/jdnCreations/aoc/2024/go"
)

func countSafeReports(reports []int) int {
	return 1
}

func checkReport(report []int, reports *int) {
  fmt.Println("------CHECKING REPORT-------")
  fmt.Println(report)
  // check if first value is greater than second value
  mode := ""

  if report[0] > report[1] {
    mode = "decreasing"
  }

  if report[0] < report[1] {
    mode = "increasing"
  }

  for i := range report[1:] {
    if report[i] > report[i+1] && mode != "decreasing" {
      fmt.Println("UNSAFE")
      return 
    }
    if report[i] < report[i+1] && mode != "increasing" {
      fmt.Println("UNSAFE")
      return
    }
    dist := utils.GetDistance(report[i], report[i+1])
    if dist > 3 || dist == 0 {
      fmt.Println("distance too large")
      return 
    }
  }
  fmt.Println("SAFE")
  *reports += 1
}

func convertStrSliceToInt(slice []string) []int {
  intSlice := make([]int, len(slice))

  for i, s := range slice {
    num, err := strconv.Atoi(s)
    if err != nil {
      fmt.Println("Error:", err)
      continue
    }
    intSlice[i] = num
  }
  return intSlice
}

func solvePartOne(scanner bufio.Scanner) int {
  safeReports := 0

	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
    intValues := convertStrSliceToInt(values)
    checkReport(intValues, &safeReports)
	}
	return safeReports 
}

func main() {
	file, err := os.Open("./day2input.txt")
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	part1 := solvePartOne(*scanner)
	fmt.Printf("Safe reports: %d\n", part1)
}