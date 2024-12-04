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



func checkReport(report []int) string {
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
      return "unsafe" 
    }
    if report[i] < report[i+1] && mode != "increasing" {
      return "unsafe"
    }
    dist := utils.GetDistance(report[i], report[i+1])
    if dist > 3 || dist == 0 {
      return "unsafe"
    }
  }
  return "safe"
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
    res := checkReport(intValues)
    if res == "safe" {
      safeReports++
    }
	}
	return safeReports 
}

func useProblemDampener(reports []int) string {
  for i := range reports {
    // run checkReport with reports missing first int
    reportMissingFirst := make([]int, len(reports))
    copy(reportMissingFirst, reports)
    removedItem := append(reportMissingFirst[:i], reportMissingFirst[i+1:]...)
    res := checkReport(removedItem)
    if res == "safe" {
      return "safe"
    }
  }
  return "unsafe"
}

func solvePartTwo (scanner bufio.Scanner) int {
  safeReports := 0

  for scanner.Scan() {
    values := strings.Fields(scanner.Text())
    intValues := convertStrSliceToInt(values)
    res := checkReport(intValues)
    if res == "safe" {
      safeReports++
    }
    if res == "unsafe" {
      // run check with problem dampener
      dampRes := useProblemDampener(intValues)
      if dampRes == "safe" {
        safeReports++
      }
    }
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
	part1 := solvePartTwo(*scanner)
	fmt.Printf("Safe reports: %d\n", part1)
}