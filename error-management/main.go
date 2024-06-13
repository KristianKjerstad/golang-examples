package main

import (
	"errors"
	"fmt"
)

func main() {

	err := errors.New("an err")
	fmt.Println(err)

	err2 := fmt.Errorf("this error wraps the first: %w", err)

	fmt.Println(err2)
	divide(1, 2)
	divide(1, 0)

	result, err := divideWithError(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func divide(dividend, divisor int) int {
	defer func() {

		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return dividend / divisor
}

func divideWithError(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("invalid divisor: cant divide by zero")
	}
	return dividend / divisor, nil
}
