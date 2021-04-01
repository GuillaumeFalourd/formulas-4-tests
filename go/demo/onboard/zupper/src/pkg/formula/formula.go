package formula

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Zupper struct {
	Email   string
	Profile string
	Pod     string
	Project string
}

func (z Zupper) Run() {
	split := strings.Split(z.Email, ".")
	name := split[0]
	surname := strings.Split(split[1], "@")[0]
	fmt.Println("Starting onboarding process for Zupper", strings.ToUpper(name+" "+surname))
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	if z.Project != "CROSS" {
		switch z.Profile {
		case "BACK-END":
			fmt.Println("Freeing access to " + z.Project + " repository from " + z.Pod + " team on GITHUB")
		case "FRONT-END":
			fmt.Println("Freeing access to " + z.Project + " repository from " + z.Pod + " team on GITLAB")
		case "SRE":
			fmt.Println("Freeing access to " + z.Project + " pipelines from " + z.Pod + " team on JENKINS")
		case "QA":
			fmt.Println("Freeing access to " + z.Project + " collection from " + z.Pod + " team on POSTMAN")
		}
	} else {
		switch z.Profile {
		case "BACK-END":
			fmt.Println("Freeing access to all repositories from " + z.Pod + " team on GITHUB")
		case "FRONT-END":
			fmt.Println("Freeing access to all repositories from " + z.Pod + " team on GITLAB")
		case "SRE":
			fmt.Println("Freeing access to all pipelines from " + z.Pod + " team on JENKINS")
		case "QA":
			fmt.Println("Freeing access to all collections from " + z.Pod + " team on POSTMAN")
		}
	}
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	log.Println("......")
	time.Sleep(time.Second * 1)
	fmt.Println("✅ Onboarding successfully executed, welcome aboard " + strings.ToUpper(name+" "+surname) + " !")
}
