package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func runCode(position int, program []int64) (int, []int64) {
	//fmt.Printf("Running position %d, which is opcode %d\n", position, program[position])
	if program[position] == 99 {
		return position, program
	}
	aLoc := program[position+1]
	bLoc := program[position+2]
	cLoc := program[position+3]
	if program[position] == 1 {
		program[cLoc] = program[aLoc] + program[bLoc]
	} else if program[position] == 2 {
		program[cLoc] = program[aLoc] * program[bLoc]
	} else {
		panic("wrong opcode")
	}
	newPosition := position + 4
	newProgram := program
	return runCode(newPosition, newProgram)
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
	var output []rune
	for {
		c, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, c)
	}
	//split comma separated ints into an int array
	var program []int64
	for _, element := range strings.Split(string(output), ",") {
		num, _ := strconv.ParseInt(element, 10, 64)
		program = append(program, num)
	}

	//ORIGINAL STEP1:
	/*
			position := 0
			position, program = runCode(position, program)
				fmt.Printf("final program:\n")
		for _, element := range program {
			fmt.Printf("%d,", element)
		}
	*/
	//STEP2:
	for n := int64(1); n <= 100; n++ {
		for v := int64(1); v <= 100; v++ {
			testProgram := make([]int64, len(program))
			copy(testProgram, program)
			testProgram[1] = n
			testProgram[2] = v
			_, testProgram = runCode(0, testProgram)
			fmt.Printf("For N: %d V: %d, position 0 ends: %d\n", n, v, testProgram[0])
			if testProgram[0] == 19690720 {
				return
			}
		}
	}

}
