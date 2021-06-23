package main

import (
	"formula/pkg/formula"
	"os"
	"strconv"
)

func main() {
	input1, _ := strconv.Atoi(os.Getenv("RIT_NUMBER_ONE"))
	input2, _ := strconv.Atoi(os.Getenv("RIT_NUMBER_TWO"))

	formula.Formula{
		NumberOne: input1,
		NumberTwo: input2,
	}.Run()
}
