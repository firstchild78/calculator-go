package selenium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const url = "http://localhost:1235"

func TestSeleniumAddValids(t *testing.T) {

	wd, sv := setUp(url)
	defer sv.Stop()
	res := calculate(wd, "4", "2", "+")
	assert.Equal(t, "6", res)

}

func TestSeleniumSubtractValids(t *testing.T) {

	wd, sv := setUp(url)
	defer sv.Stop()
	res := calculate(wd, "4", "2", "-")
	assert.Equal(t, "2", res)

}

func TestSeleniumMultiplyValids(t *testing.T) {

	wd, sv := setUp(url)
	defer sv.Stop()
	res := calculate(wd, "4", "2", "*")
	assert.Equal(t, "8", res)

}

func TestSeleniumDivideValids(t *testing.T) {

	wd, sv := setUp(url)
	defer sv.Stop()
	res := calculate(wd, "4", "2", "/")
	assert.Equal(t, "2", res)

}

func TestSeleniumDivideByZero(t *testing.T) {

	wd, sv := setUp(url)
	defer sv.Stop()
	res := calculate(wd, "4", "0", "/")
	assert.Equal(t, "0", res)

}
