// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"project/pkg/tpl"

	"github.com/fatih/color"
)

const dirFormat = "%s/%s"

type Input struct {
	Username        string
	ProjectName     string
	ProjectLocation string
}

func (in Input) Path() string {
	return fmt.Sprintf(dirFormat, in.ProjectLocation, in.ProjectName)
}

func (in Input) Run() {
	if err := CreateDirIfNotExists(in.Path(), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create project: %q, error: %q", in.Path(), err.Error()))
		os.Exit(1)
	}

	ymlworkflow := fmt.Sprintf(in.Path())
	if err := CreateFileIfNotExist(ymlworkflow, []byte("")); err != nil {
		color.Red(fmt.Sprintf("failed to create project: %q, error: %q", in.Path(), err.Error()))
		os.Exit(1)
	}

	t := template.Must(template.New("workflow").Parse(tpl.Workflow))
	bfile, err := os.OpenFile(ymlworkflow, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		color.Red(fmt.Sprintf("failed to create project: %q, error: %q", in.Path(), err.Error()))
		os.Exit(1)
	}
	defer bfile.Close()
	err = t.Execute(bfile, in)
	if err != nil {
		color.Red(fmt.Sprintf("failed to create project: %q, error: %q", in.Path(), err.Error()))
		os.Exit(1)
	}
}

func CreateDirIfNotExists(dir string, perm os.FileMode) error {
	if IsNotExist(dir) {
		if err := os.MkdirAll(dir, perm); err != nil {
			return fmt.Errorf("failed to create directory: %q, error: %q", dir, err)
		}
	}
	return nil
}

func CreateFileIfNotExist(file string, content []byte) error {
	if IsNotExist(file) {
		if err := ioutil.WriteFile(file, content, 0644); err != nil {
			return err
		}
	}
	return nil
}

func IsNotExist(name string) bool {
	_, err := os.Stat(name)
	return err != nil && os.IsNotExist(err)
}
