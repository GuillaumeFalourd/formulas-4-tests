package formula

import (
	"fmt"
	"strings"
	"time"
)

type Function struct {
	Profile string
	Need    string
	Ritchie bool
}

func (f Function) Run() {
	fmt.Println("Hi " + strings.ToLower(f.Profile) + " professional looking for answers...")
	time.Sleep(time.Second * 2)
	if f.Ritchie {
		fmt.Sprintf("You already know Ritchie is the best tool for you !")
	} else {
		fmt.Println("If your need " + strings.ToLower(f.Need) + "...")
		time.Sleep(time.Second * 2)
		fmt.Println("You have to test Ritchie now, this tool has been made for you !")
	}
}
