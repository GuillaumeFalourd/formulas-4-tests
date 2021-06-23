// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"os"
	"project/pkg/formula"
)

const (
	GITHUB_TOKEN = " ${{ secrets.GITHUB_TOKEN }} "
	REPO_OWNER   = " ${{ github.repository_owner }} "
	ACCESS_TOKEN = " ${{ secrets.ACCESS_TOKEN }} "
)

func main() {
	os.Setenv("GITHUB_TOKEN", GITHUB_TOKEN)
	os.Setenv("REPO_OWNER", REPO_OWNER)
	os.Setenv("ACCESS_TOKEN", ACCESS_TOKEN)

	username := os.Getenv("USERNAME")
	projectName := os.Getenv("PROJECT_NAME")
	projectLocation := os.Getenv("PROJECT_LOCATION")

	formula.Input{
		Username:        username,
		ProjectName:     projectName,
		ProjectLocation: projectLocation,
	}.Run()
}
