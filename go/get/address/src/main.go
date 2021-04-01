package main

import (
	"formula/pkg/formula"
	"log"
	"os"
	"regexp"
)

func main() {
	cep := os.Getenv("CEP")

	if isValidCEP(cep) {
		formula.GetAddress(cep)
	}
}

func isValidCEP(cep string) bool {
	cepMatchDigit, _ := regexp.MatchString(`^\d{8}$`, cep)
	cepMatchPattern, _ := regexp.MatchString(`^\d{5}-\d{3}$`, cep)

	if !cepMatchDigit && !cepMatchPattern {
		log.Printf("The value informed : %s, is not a valid CEP", cep)
		return false
	}

	return true
}
