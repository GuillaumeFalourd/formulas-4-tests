package formula

import (
	"bytes"
	"encoding/json"
	"log"
)

type Formula struct {
	Username string
	Token    string
}

func (h Formula) Run() {
	response, err := json.Marshal(h)
	if err != nil {
		log.Printf("Error", h)
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, response, "", "\t")
	if error != nil {
		log.Println("JSON parse error:", error)
	}

	log.Println("Github Credentials:", string(prettyJSON.Bytes()))
}
