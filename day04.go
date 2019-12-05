package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func getDigit(num int64, place int64) int64 {
	return (num / int64(math.Pow(10, float64(place)))) % 10
}
func hasAdjacent(num int64) bool {
	var x, y int64
	adjCount := int64(1)
	x = getDigit(num, 0)
	for i := int64(1); i < 6; i++ {
		y = getDigit(num, i)
		if x == y {
			adjCount++ //step 2
			// return true
		} else { //step 2
			if adjCount == 2 {
				return true
			}
			adjCount = 1
		}
		x = y
	}
	if adjCount == 2 { //step 2
		return true
	}
	return false
}
func isAscending(num int64) bool {
	var x, y int64
	x = getDigit(num, 5)
	for i := int64(4); i >= 0; i-- {
		y = getDigit(num, i)
		if y < x {
			return false
		}
		x = y
	}
	return true
}

func main() {
	/* BOILERPLATE FOR READING INPUTS */
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("Pipe input data.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	if err != nil && err == io.EOF {
		panic(err)
	}
	//Day 04
	rangeValues := strings.Split(string(input), "-")
	a, _ := strconv.ParseInt(rangeValues[0], 10, 64)
	b, _ := strconv.ParseInt(rangeValues[1], 10, 64)
	count := int64(0)
	for i := a; i <= b; i++ {
		if hasAdjacent(i) && isAscending(i) {
			fmt.Printf("Found %d\n", i)
			count++
		}
	}
	fmt.Printf("Count: %d\n", count)
}
