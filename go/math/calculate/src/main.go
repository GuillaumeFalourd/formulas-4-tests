package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	numberOne := os.Getenv("RIT_NUMBER_ONE")
	numberTwo := os.Getenv("RIT_NUMBER_TWO")
	operation := os.Getenv("RIT_OPERATION")

	formula.Formula{
		NumberOne: numberOne,
		NumberTwo: numberTwo,
		Operation: operation,
	}.Run()
}
