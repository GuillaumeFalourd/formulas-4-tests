package formula

import (
	"fmt"
	"log"
	"os/exec"
)

type Formula struct {
	NumberOne string
	NumberTwo string
	Operation string
}

func (h Formula) Run() {
	cmdLine := exec.Command("", "")

	switch h.Operation {
	case "sum":
		cmdLine = exec.Command(
			"rit",
			"go",
			"math",
			"sum",
			"numbers",
			fmt.Sprintf("--rit_number_one=%s", h.NumberOne),
			fmt.Sprintf("--rit_number_two=%s", h.NumberTwo),
		)
	case "multiplication":
		cmdLine = exec.Command(
			"rit",
			"go",
			"math",
			"multiply",
			"numbers",
			fmt.Sprintf("--rit_number_one=%s", h.NumberOne),
			fmt.Sprintf("--rit_number_two=%s", h.NumberTwo),
		)
	default:
		fmt.Print("Unexpected operation type:", h.Operation)
	}

	out, err := cmdLine.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Printf(string(out))

}
