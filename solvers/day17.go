package solvers

import (
	"aoc-2024/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	common.RegisterSolver(17, SolveDay17)
}

type ExecutionStep struct {
	outputType string
	value      int
}

func getComboVal(combo int, registers []int) int {
	if combo == 7 {
		panic("Invalid combo value")
	}
	if combo < 4 {
		return combo
	}
	return registers[combo-4]
}

func execute(registers []int, opcode int, operand int) ExecutionStep {
	var outputType string
	var value int
	switch opcode {
	case 0:
		num := registers[0]
		comboVal := getComboVal(operand, registers)
		denom := math.Pow(2, float64(comboVal))
		result := num / int(denom)
		registers[0] = int(result)
		break
	case 1:
		num := registers[1]
		result := num ^ operand
		registers[1] = result
		break
	case 2:
		num := getComboVal(operand, registers)
		result := num % 8
		registers[1] = result
		break
	case 3:
		if registers[0] == 0 {
			// noop
		} else {
			value = operand
			outputType = "jump"
		}
		break
	case 4:
		num1 := registers[1]
		num2 := registers[2]
		result := num1 ^ num2
		registers[1] = result
		break
	case 5:
		num := getComboVal(operand, registers)
		value = num % 8
		outputType = "out"
		break
	case 6:
		num := registers[0]
		comboVal := getComboVal(operand, registers)
		denom := math.Pow(2, float64(comboVal))
		result := num / int(denom)
		registers[1] = int(result)
		break
	case 7:
		num := registers[0]
		comboVal := getComboVal(operand, registers)
		denom := math.Pow(2, float64(comboVal))
		result := num / int(denom)
		registers[2] = int(result)
		break
	}

	return ExecutionStep{outputType: outputType, value: value}
}

func SolveDay17(input string) string {
	inputs := strings.Split(input, "\n\n")
	registers := []int{0, 0, 0} // A B C
	registerInputs := strings.Split(inputs[0], "\n")
	for i := range registerInputs {
		reg := strings.Split(registerInputs[i], ": ")
		num, err := strconv.Atoi(reg[1])
		if err == nil {
			registers[i] = num
		}
	}

	executionSteps := []int{}
	programInput := strings.Split(strings.Split(inputs[1], ": ")[1], ",")
	for i := range programInput {
		num, err := strconv.Atoi(programInput[i])
		if err == nil {
			executionSteps = append(executionSteps, num)
		}
	}

	step := 0
	outs := []string{}
	for step < len(executionSteps)-1 {
		res := execute(registers, executionSteps[step], executionSteps[step+1])
		if res.outputType == "jump" {
			step = res.value
		} else {
			step += 2
		}

		if res.outputType == "out" {
			outs = append(outs, strconv.Itoa(res.value))
		}
	}

	result1 := strings.Join(outs, ",")

	result2 := 0
	// 1,4,6,1,6,4,3,0,3
	return fmt.Sprintf("%s\n%d", result1, result2)
}
