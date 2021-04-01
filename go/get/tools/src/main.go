// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	name := os.Getenv("RIT_NAME")
	profile := os.Getenv("RIT_PROFILE")
	tool := os.Getenv("RIT_TOOL")

	formula.Formula{
		Name:    name,
		Profile: profile,
		Tool:    tool,
	}.Run(os.Stdout)
}
