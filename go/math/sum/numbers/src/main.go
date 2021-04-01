package main

import (
	"formula/pkg/formula"
	"os"
	"strconv"
)

func main() {
	input1, _ := strconv.Atoi(os.Getenv("NUMBER_ONE"))
	input2, _ := strconv.Atoi(os.Getenv("NUMBER_TWO"))

	formula.Formula{
		NumberOne: input1,
		NumberTwo: input2,
	}.Run()
}
