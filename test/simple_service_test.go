package test

import (
	"belajar_golang_api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	service, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, service)
}
func TestSimpleServiceSuccess(t *testing.T) {
	service, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, service)
}
