package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var left = 4
var right = 2

func TestAddValids(t *testing.T) {
	expected := 6
	actual := Add(left, right)

	assert.Equal(t, expected, actual)
}

func TestSubtractValids(t *testing.T) {
	expected := 2
	actual := Subtract(left, right)

	assert.Equal(t, expected, actual)
}

func TestMultiplyValids(t *testing.T) {
	expected := 8
	actual := Multiply(left, right)

	assert.Equal(t, expected, actual)
}

func TestDivisionTableTest(t *testing.T) {
	v := []struct {
		left     int
		right    int
		expected int
	}{
		{4, 2, 2},
		{10, 5, 2},
		{2, 0, 0},
		{-9, 3, -3},
	}
	for _, v := range v {
		assert.Equal(t, v.expected, Divide(v.left, v.right))
	}
}
