package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	assert := assert.New(t)
	existKey := "testing exists key"
	existKeyvalue := "testing exists value"
	nonexistentKey := "testing not exists key"
	os.Setenv(existKey, existKeyvalue)

	assert.Equal(GetEnv(existKey, "test"), existKeyvalue)
	assert.Equal(GetEnv(nonexistentKey, "test"), "test")
}
