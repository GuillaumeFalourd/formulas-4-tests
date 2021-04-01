package formula

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	statsURL   = "https://corona-stats.online"
	topURL     = "%s?top=%s"
	countryURL = "%s/%s"
	graphURL   = "%s/%s/graph"
	top        = "top N"
	graph      = "graph by country"
	stats      = "stats by country"
)

//Input represents the input values of the formula
type Input struct {
	Search string
	Value  string
}

//Run runs the formula logic
func (in Input) Run() {
	var url string
	switch in.Search {
	case top:
		url = fmt.Sprintf(topURL, statsURL, in.Value)
	case graph:
		url = fmt.Sprintf(graphURL, statsURL, in.Value)
	case stats:
		url = fmt.Sprintf(countryURL, statsURL, in.Value)
	default:
		log.Fatalln("Search type not recognized. Please use (stats, graph, top)")
	}

	cmd := exec.Command("curl", "-s", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

}
