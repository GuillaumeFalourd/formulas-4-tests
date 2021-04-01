package formula

import (
	"fmt"
	"os"
	"os/exec"
)

type Formula struct {
	NumberOne string
	NumberTwo string
	Operation string
}

func (h Formula) Run() {

	json := fmt.Sprintf("'{\"number_one\":\"" + h.NumberOne + "\", \"number_two\":\"" + h.NumberTwo + "\"}'")

	switch h.Operation {
	case "sum":
		cmdLine := "rit go math sum numbers --stdin"
		stdinExec(json, cmdLine)
	case "multiplication":
		cmdLine := "rit go math multiply numbers --stdin"
		stdinExec(json, cmdLine)
	default:
		fmt.Printf("Unexpected operation type:", h.Operation)
	}
}

func stdinExec(json string, cmdLine string) {

	echo := "echo"
	arg0 := json
	arg1 := "|"
	arg2 := cmdLine

	cmdStdin := &exec.Cmd{
		Path:   os.Getenv("CURRENT_PWD"),
		Args:   []string{"bash", echo, arg0, arg1, arg2},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	fmt.Println(cmdStdin.String())

	err := cmdStdin.Run() //TODO: Make it work!
	if err != nil {
		fmt.Println(err)
	}

	//cmd := exec.Command("bash", echo, arg0, arg1, arg2)
	//cmd.Dir = os.Getenv("CURRENT_PWD")
	//var out bytes.Buffer
	//var stderr bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = &stderr
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	//}
	//fmt.Println("Result: " + out.String())

	//cmd := exec.Command(cmdLine)
	//err := cmd.Run()
	//if err != nil {
	//	println("Error", err)
	//}

	//stdout, err := cmd.Output()

	//if err != nil {
	//	println("Error:", err)
	//	return
	//}

	//println("Output:", string(stdout))
}
