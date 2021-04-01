package formula

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const notFound = "Didn't find any address for the CEP informed : %s"

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

func GetAddress(cep string) {
	response, err := http.Get(
		fmt.Sprintf(
			"http://viacep.com.br/ws/%s/json/",
			strings.ReplaceAll(cep, "-", ""),
		),
	)
	if err != nil {
		return
	}

	var address Address
	err = json.NewDecoder(response.Body).Decode(&address)
	if err != nil {
		return
	}

	if address.Cep == "" {
		log.Printf(notFound, cep)
		return
	}

	addressJson, _ := getAddressJson(address, cep)

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, addressJson, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}

	log.Println("Address:", string(prettyJSON.Bytes()))
	return
}

func getAddressJson(address Address, cep string) ([]byte, error) {
	res, err := json.Marshal(address)
	if err != nil {
		log.Printf(notFound, cep)
		return []byte{}, err
	}

	if res == nil {
		log.Printf(notFound, cep)
		return []byte{}, err
	}

	return res, nil
}
