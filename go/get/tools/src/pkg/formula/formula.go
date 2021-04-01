// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io"

	"github.com/gookit/color"
)

type Formula struct {
	Name    string
	Profile string
	Tool    string
}

func (f Formula) Run(writer io.Writer) {
	var result string

	result += color.FgGreen.Render(fmt.Sprintf("Welcome to %s.\n", f.Name))

	result += color.FgYellow.Render(fmt.Sprintf("He is our new %s professional.\n", f.Profile))

	result += color.FgCyan.Render(fmt.Sprintf("Ritchie will install %s for him easily according to his OS.\n", f.Tool))

	if _, err := fmt.Fprintf(writer, result); err != nil {
		panic(err)
	}
}
