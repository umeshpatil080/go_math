package go_math

import "errors"

func Add(args ...int) int {
    result := 0
    for _, val := range args {
	result += val
    }
    return result
}

func Sub(a, b int) int {
	return a - b
}

func Mul(args ...int) int {
    result := 1
    for _, val := range args {
	result *= val
    }
    return result
}

func Div(a, b int) (int, error) {
	if b == 0 {
	    return 0, errors.New("Divide by Zero is not allowed")
	}
	return a / b, nil
}
