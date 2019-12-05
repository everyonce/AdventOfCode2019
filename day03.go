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

type point struct {
	x int64
	y int64
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
	var lines [][]byte
	for {
		c, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, c)
	}

	//Day 03
	var origin = point{x: 0, y: 0}
	var dirPoints = make(map[byte]point)
	dirPoints['R'] = point{x: 1, y: 0}
	dirPoints['L'] = point{x: -1, y: 0}
	dirPoints['U'] = point{x: 0, y: 1}
	dirPoints['D'] = point{x: 0, y: -1}
	var prev point
	var cablePaths = make(map[int][]point)
	for index, line := range lines {
		prev = origin
		cablePaths[index] = append(cablePaths[index], prev)
		for _, instruction := range strings.Split(string(line), ",") {
			val, _ := strconv.ParseInt(instruction[1:], 10, 64)
			for i := int64(1); i <= val; i++ {
				var newPoint = point{x: prev.x + dirPoints[instruction[0]].x,
					y: prev.y + dirPoints[instruction[0]].y}
				cablePaths[index] = append(cablePaths[index], newPoint)
				prev = newPoint
			}
		}
	}
	mDistanceMin := float64(0)
	mStepsMin := float64(0) //step 2
	first := true
	for i, p1 := range cablePaths[0] {
		for j, p2 := range cablePaths[1] {
			if p1.x == p2.x && p1.y == p2.y {
				fmt.Printf("intersection %d,%d to %d,%d: ", p1.x, p1.y, p2.x, p2.y)
				mD := math.Abs(float64(p1.x)) + math.Abs(float64(p1.y))
				mS := float64(i + j) //step 2
				if first && mD > 0 {
					first = false
					mDistanceMin = mD
					mStepsMin = mS //step 2
				} else {
					mDistanceMin = math.Min(mDistanceMin, mD)
					mStepsMin = math.Min(mStepsMin, mS) //step 2
				}
				fmt.Printf("md %f ms %f\n", mD, mS)

			}
		}
	}
	fmt.Printf("Minimum mDistance: %f\n", mDistanceMin)
	fmt.Printf("Minimum mSteps: %f\n", mStepsMin)
}
