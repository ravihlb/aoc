package main

import (
	"testing"
)

func TestPointCalc(t *testing.T) {
	const filepath string = "./example.input"
	const expected int = 13

	actual := solve(filepath)

	if actual != expected {
		t.Errorf("answer = %d, expected %d", actual, expected)
	}
}
