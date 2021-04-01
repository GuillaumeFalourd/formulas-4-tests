package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	email := os.Getenv("ZUPPER_EMAIL")
	profile := os.Getenv("ZUPPER_PROFILE")
	pod := os.Getenv("ZUPPER_POD")
	project := os.Getenv("ZUPPER_PROJECT")

	formula.Zupper{
		Email:   email,
		Profile: profile,
		Pod:     pod,
		Project: project,
	}.Run()
}
