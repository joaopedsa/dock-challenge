package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCPF(t *testing.T) {
	assert := assert.New(t)

	correctCPF := "09662672907"
	assert.True(ValidateCPF(correctCPF))
	incorrectCPF := "09662672902"
	assert.False(ValidateCPF(incorrectCPF))
}

func TestValidateCNPJ(t *testing.T) {
	assert := assert.New(t)

	correctCNPJ := "55335356000178"
	assert.True(ValidateCNPJ(correctCNPJ))
	incorrectCNPJ := "13246569000182"
	assert.False(ValidateCNPJ(incorrectCNPJ))
}

func TestRemoveNonNumbersCPF(t *testing.T) {
	assert := assert.New(t)

	cpf := "096.626.729-07"
	expectedCPF := "09662672907"

	assert.Equal(RemoveNonNumbersCPF(cpf), expectedCPF)
}
