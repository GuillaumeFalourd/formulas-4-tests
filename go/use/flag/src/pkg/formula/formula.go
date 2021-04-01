// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

type Formula struct {
	Text     string
	List     string
	Boolean  bool
	Password string
	Mode     string
}

// INPUT FLAG: rit demo hello-world rit_input_text='f.Text' rit_input_boolean=f.Boolean rit_input_list='f.List' rit_input_password='f.Password'
// STDIN FLAG: echo '{"rit_input_text":"Dennis", "rit_input_boolean":"true", "rit_input_list":"everything", "rit_input_password":"Ritchie"}' | rit demo hello-world --stdin

func (f Formula) Run(writer io.Writer) {

	cmd := exec.Command("", "")

	if f.Mode == "input flag" {
		cmd = exec.Command(
			"rit",
			"demo",
			"hello-world",
			fmt.Sprintf("--rit_input_text=%s", f.Text),
			fmt.Sprintf("--rit_input_list=%s", f.List),
			fmt.Sprintf("--rit_input_boolean=%v", f.Boolean),
			fmt.Sprintf("--rit_input_password=%s", f.Password),
		)

	} else if f.Mode == "stdin flag" {
		cmd = exec.Command(
			"bash",
			"-c",
			fmt.Sprintf(
				"echo '{\"rit_input_text\":\"%s\", \"rit_input_boolean\":\"%v\", \"rit_input_list\":\"%s\", \"rit_input_password\":\"%s\"}' | rit demo hello-world --stdin",
				f.Text,
				f.Boolean,
				f.List,
				f.Password,
			),
		)

	} else {
		log.Fatalf("Unexpected Mode: %s\n", f.Mode)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Printf(string(out))
}
