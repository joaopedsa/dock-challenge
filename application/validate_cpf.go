package application

import "strconv"

func ValidateCPF(CPF string) bool {
	var soma int32
	var resto int32
	if CPF == "00000000000" {
		return false
	}
	for i := 1; i <= 9; i++ {
		digitCPF, err := strconv.ParseInt(CPF[i-1:i], 10, 32)
		if err != nil {
			return false
		}
		soma = soma + int32(digitCPF)
	}
	resto = (soma * 10) % 11
	if resto == 10 || resto == 11 {
		resto = 0
	}

	digitCPF, err := strconv.ParseInt(CPF[9:10], 10, 32)
	if err != nil {
		return false
	}
	if resto != int32(digitCPF) {
		return false
	}

	for i := 1; i <= 10; i++ {
		digitCPF, err := strconv.ParseInt(CPF[i-1:i], 10, 32)
		if err != nil {
			return false
		}
		soma = soma + int32(digitCPF*int64(12-i))
	}
	resto = (soma * 10) % 11
	if resto == 10 || resto == 11 {
		resto = 0
	}

	digitCPF, err = strconv.ParseInt(CPF[10:11], 10, 32)
	if err != nil {
		return false
	}
	if resto != int32(digitCPF) {
		return false
	}

	return true
}
