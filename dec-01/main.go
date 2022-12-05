package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// file, err := os.Open("dec-01/calories.txt")
	// if err != nil {
	// panic(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	scanner := initScanner("dec-01/calories.txt")

	totals := []int{}
	temp := 0

	for scanner.Scan() {
		res, _ := strconv.Atoi(scanner.Text())
		switch res {
		case 0:
			totals = append(totals, temp)
			temp = 0
		default:
			temp += res
		}
	}

	answer := 0
	sort.Ints(totals)
	for i := 0; i < 3; i++ {
		answer += totals[len(totals)-i-1]
	}
	fmt.Println(answer)
	fmt.Println(totals)
}

func first() {
	temp := 0
	highest := 0

	scanner := initScanner("dec-01/calories.txt")
	for scanner.Scan() {
		res, _ := strconv.Atoi(scanner.Text())
		switch res {
		case 0:
			if temp > highest {
				highest = temp
			}
			temp = 0
		default:
			temp += res
		}
	}
	fmt.Println(highest)
}

func initScanner(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// defer file.Close()

	return bufio.NewScanner(file)
}
