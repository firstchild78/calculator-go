package main

//Add two numbers
func Add(left int, right int) int {
	return left + right
}

//Subtract two numbers
func Subtract(left int, right int) int {
	return left - right
}

//Multiply two numbers
func Multiply(left int, right int) int {
	return left * right
}

//Divide two numbers
func Divide(left int, right int) int {
	if right == 0 {
		return 0
	}
	return left / right
}
