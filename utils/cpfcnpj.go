package utils

import (
	"regexp"

	"github.com/klassmann/cpfcnpj"
)

func ValidateCPF(CPF string) bool {
	return cpfcnpj.ValidateCPF(CPF)
}

func ValidateCNPJ(CNPJ string) bool {
	return cpfcnpj.ValidateCNPJ(CNPJ)
}

func RemoveNonNumbersCPF(CPF string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	return nonAlphanumericRegex.ReplaceAllString(CPF, "")
}
