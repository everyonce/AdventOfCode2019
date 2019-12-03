package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func getFuel(mass float64) int64 {
	//ORIGINAL: return int64(math.Floor(mass/3) - 2)
	fuel := int64(math.Floor(mass/3) - 2)
	if fuel <= 0 {
		return 0
	}
	return fuel + getFuel(float64(fuel))
}

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("Pipe input data.")
		fmt.Println("Usage: day1a | mydata")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output [][]byte

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, line)
	}
	total := int64(0)
	for j := 0; j < len(output); j++ {
		mass, err := strconv.ParseFloat(string(output[j]), 64)
		if err != nil {
			panic(err)
		}
		total += getFuel(mass)
	}
	fmt.Printf("%d total fuel\n", total)

}
