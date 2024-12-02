package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFileLists(filename string) ([]int, []int, error) {
	var list1, list2 []int
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineNumber int
	lineNumber = 0
	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		// fmt.Println(numbers)
		// fmt.Println(len(numbers))
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("wrong file format in line: %d", lineNumber)
		}

		number1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, nil, fmt.Errorf("wrong number format in line %d", lineNumber)
		}
		number2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, nil, fmt.Errorf("wrong number format in line %d", lineNumber)
		}

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}

	return list1, list2, nil
}

func readFileListMap(filename string) ([]int, map[int]int, error) {
	var list []int
	freq := make(map[int]int)
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineNumber int
	lineNumber = 0
	for scanner.Scan() {
		lineNumber += 1
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		// fmt.Println(numbers)
		// fmt.Println(len(numbers))
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("wrong file format in line: %d", lineNumber)
		}

		number1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return nil, nil, fmt.Errorf("wrong number format in line %d", lineNumber)
		}
		number2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return nil, nil, fmt.Errorf("wrong number format in line %d", lineNumber)
		}

		list = append(list, number1)

		freq[number2] += 1
	}

	return list, freq, nil
}

func absDiff(number1, number2 int) int {
	result := number1 - number2
	if result < 0 {
		return result * -1
	}
	return result
}

func part1(result chan int) {
	defer close(result)
	list1, list2, err := readFileLists("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	sort.Ints(list1)
	sort.Ints(list2)

	var totalDifference int
	totalDifference = 0
	for i := 0; i < len(list1); i++ {
		difference := absDiff(list1[i], list2[i])
		totalDifference += difference
	}
	result <- totalDifference
	return
}

func part2(result chan int) {
	defer close(result)
	numbers, freq, err := readFileListMap("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var similarityScore int
	similarityScore = 0
	for i := 0; i < len(numbers); i++ {
		similarityScore += numbers[i] * freq[numbers[i]]
	}
	result <- similarityScore
	return
}

func main() {
	totalDifferenceResult := make(chan int)
	go part1(totalDifferenceResult)
	similarityScoreResult := make(chan int)
	go part2(similarityScoreResult)
	fmt.Printf("Solution for Part 1:\nThe Total Difference is %d\n\n", <-totalDifferenceResult)
	fmt.Printf("Solution for Part 2:\nThe Similarity Score is %d\n", <-similarityScoreResult)
	return
}
