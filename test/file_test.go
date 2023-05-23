package test

import (
	"belajar_golang_api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("test")

	assert.NotNil(t, connection)
	cleanup()
}
