package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	input1 := os.Getenv("NUMBER_ONE")
	input2 := os.Getenv("NUMBER_TWO")
	input3 := os.Getenv("OPERATION")

	formula.Formula{
		NumberOne: input1,
		NumberTwo: input2,
		Operation: input3,
	}.Run()
}
