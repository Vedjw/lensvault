package main

import (
	"errors"
	"fmt"
)

var errVar = errors.New("not found")

func A() error {
	return errVar
}

func B() error {
	err := A()

	return fmt.Errorf("B: %w", err)
}

func main() {
	err := B()

	fmt.Print(err)
}
