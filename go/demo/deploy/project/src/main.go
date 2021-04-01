package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	name := os.Getenv("PROJECT_NAME")
	tag := os.Getenv("RELEASE_TAG")
	env := os.Getenv("ENVIRONMENT")
	dt := os.Getenv("DEPLOYMENT_TYPE")

	formula.Deploy{
		Name: name,
		Tag:  tag,
		Env:  env,
		Type: dt,
	}.Run()
}
