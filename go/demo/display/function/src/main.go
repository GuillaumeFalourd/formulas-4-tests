package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	var ritchie bool
	profile := os.Getenv("PROFILE")
	need := os.Getenv("NEED")
	boolean := os.Getenv("RITCHIE")

	if boolean == "yes" {
		ritchie = true
	}

	formula.Function{
		Profile: profile,
		Need:    need,
		Ritchie: ritchie,
	}.Run()
}
