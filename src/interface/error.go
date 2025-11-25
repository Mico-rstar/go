package main

import "fmt"

type DivisionError struct {
	Dividend float64
	Divisor  float64
	Message  string
}

func (e DivisionError) Error() string {
	return fmt.Sprintf("除法错误: %.2f / %.2f - %s",
		e.Dividend, e.Divisor, e.Message)
}

func New(a, b float64) DivisionError {
	return DivisionError{
		Dividend: a,
		Divisor:  b,
		Message:  "除数不为0",
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, New(a, b)
	}
	return a / b, nil
}

func ErrorTest() {

}
