package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	search := os.Getenv("SEARCH")
	value := os.Getenv("VALUE")

	in := formula.Input{
		Search: search,
		Value:  value,
	}
	in.Run()
}
