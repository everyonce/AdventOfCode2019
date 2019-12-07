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

func quickConvertSlice(slice []int64) []string {
	var p []string
	for _, i := range slice {
		p = append(p, string(i))
	}
	return p
}

func inputValue() int64 {
	return 5 //was 1 for step 1
}
func outputValue(x int64) {
	fmt.Printf(">>>>>> OUTPUT: %d <<<<<< \n", x)
}
func getDigit(num int64, place int64) int64 {
	//log.Printf("getDigit %d, %d = %d\n", num, place, (num/int64(math.Pow(10, float64(place-1))))%10)
	return (num / int64(math.Pow(10, float64(place-1)))) % 10
}
func parameterPosition(fullInstruction int64, parameter int64, instructionPosition int64, program []int64) int64 {
	//log.Printf("finding Param position for instruction: %d, param# %d, instructionPosition %d\n", fullInstruction, parameter, instructionPosition)
	//log.Printf("parameter mode digit is: %d\n", getDigit(fullInstruction, parameter+2))
	if getDigit(fullInstruction, parameter+2) == 1 {
		//	log.Println("param should be immediate mode, use position of actual parameter: ", instructionPosition+parameter)
		return instructionPosition + parameter
	}
	//log.Println("param should be position mode, use position it points to: ", program[instructionPosition+parameter])
	return program[instructionPosition+parameter]
}
func runCode(position int64, program []int64) (int64, []int64) {
	fullInstruction := program[position]
	opCode := fullInstruction % 100
	var positionChange int64

	//quick summary debug
	/*
		p := 0
		fmt.Printf("Current program area (%d-%d): ", position, position+5)
		for _, i := range program[position:] {
			fmt.Printf("%d, ", i)
			p++
			if p > 5 {
				break
			}

		}
		fmt.Printf("\n")
		//r := quickConvertSlice(program[int64(math.Max(float64(position), float64(2)))-2 : position+5])

		fmt.Printf("Running position %d, which is opcode %d\n", position, opCode)
	*/
	if opCode == 99 {
		return position, program
	} else if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
		positionChange = 4
		aVal := program[parameterPosition(fullInstruction, 1, position, program)]
		bVal := program[parameterPosition(fullInstruction, 2, position, program)]
		cLoc := program[position+3]
		//fmt.Printf("DEBUG: opCode %d, aVal %d, bVal %d, cLoc %d\n", opCode, aVal, bVal, cLoc)
		if opCode == 1 {
			program[cLoc] = aVal + bVal
		} else if opCode == 2 {
			program[cLoc] = aVal * bVal
		} else if opCode == 7 {
			if aVal < bVal {
				program[cLoc] = 1
			} else {
				program[cLoc] = 0
			}
		} else if opCode == 8 {
			if aVal == bVal {
				program[cLoc] = 1
			} else {
				program[cLoc] = 0
			}
		}
	} else if opCode == 3 || opCode == 4 {
		positionChange = 2
		if opCode == 3 {
			cLoc := program[position+1]
			inputVal := inputValue()
			program[cLoc] = inputVal
			//fmt.Println("Changing position ", cLoc, " to ", inputVal)
		}
		if opCode == 4 {
			aVal := program[parameterPosition(fullInstruction, 1, position, program)]
			//fmt.Println("outputting position ", parameterPosition(fullInstruction, 1, position, program))
			outputValue(aVal)
		}
		// ADDED FOR STEP 2
	} else if opCode == 5 || opCode == 6 {
		positionChange = 3
		if opCode == 5 { //jump-if-true
			aVal := program[parameterPosition(fullInstruction, 1, position, program)]
			bVal := program[parameterPosition(fullInstruction, 2, position, program)]
			//log.Printf("JUMP IF TRUE: %d, %d\n", aVal, bVal)
			if aVal != 0 {
				return runCode(bVal, program)
			}
		}
		if opCode == 6 { //jump-if-false
			aVal := program[parameterPosition(fullInstruction, 1, position, program)]
			bVal := program[parameterPosition(fullInstruction, 2, position, program)]
			//log.Printf("JUMP IF FALSE: %d, %d\n", aVal, bVal)
			if aVal == 0 {
				return runCode(bVal, program)
			}
		}
	} else {
		panic("invalid opcode")
	}
	newPosition := position + positionChange
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
	_, program = runCode(0, program)
	//fmt.Printf("For N: %d V: %d, position 0 ends: %d\n", n, v, testProgram[0])
	/*
		for n := int64(1); n <= 100; n++ {
			for v := int64(1); v <= 100; v++ {
				testProgram := make([]int64, len(program))
				copy(testProgram, program)
				testProgram[1] = n
				testProgram[2] = v


				if testProgram[0] == 19690720 {
					return
				}
			}
		}
	*/
}
