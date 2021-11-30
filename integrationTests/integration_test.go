package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUp(left int, right int, op string) operation {
	return operation{left, right, op}
}

func TestServicesAddValids(t *testing.T) {
	op := setUp(1, 2, "+")
	res, statusCode := sendRequest(op)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, 3, res)

}

func TestServicesSubtractValids(t *testing.T) {
	op := setUp(4, 2, "-")
	res, statusCode := sendRequest(op)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, 2, res)
}

func TestServicesMultiplyValids(t *testing.T) {
	op := setUp(4, 2, "*")
	res, statusCode := sendRequest(op)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, 8, res)
}

func TestServicesDivideValids(t *testing.T) {
	op := setUp(4, 2, "/")
	res, statusCode := sendRequest(op)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, 2, res)
}

func TestServicesDivideByZero(t *testing.T) {
	op := setUp(4, 0, "/")
	res, statusCode := sendRequest(op)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, 0, res)
}
