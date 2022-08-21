package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckCPF(t *testing.T) {
	assert := assert.New(t)

	correctCPF := "09662672907"
	assert.True(ValidateCPF(correctCPF))
	incorrectCPF := "00000000000"
	assert.False(ValidateCPF(incorrectCPF))
}
