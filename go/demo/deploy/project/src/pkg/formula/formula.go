package formula

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Deploy struct {
	Name string
	Tag  string
	Env  string
	Type string
}

func (d Deploy) Run() {
	fmt.Println("Packaging %s project", strings.ToUpper(d.Name))
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	fmt.Printf("Fetching release tag %s", d.Tag)
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	fmt.Printf("Preparing " + strings.ToUpper(d.Type) + " deployment in " + strings.ToUpper(d.Env) + " environment.")
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	fmt.Printf("The " + strings.ToUpper(d.Type) + " deployment of " + strings.ToUpper(d.Name) + " project (tag " + d.Tag + ") in " + strings.ToUpper(d.Env) + " has been executed successfully !")
}
