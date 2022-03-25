package main

import (
	"testing"
)

func TestGrahamFormula(t *testing.T) {
	expected := Result{highGrahamFormulaNumber: 140.29594272076372, highConservativeGrahamFormulaNumber: 93.25059665871122, lowGrahamFormulaNumber: 323.436754176611, lowConservativeGrahamFormulaNumber: 184.82100238663486}
	value := GrahamFormula(15.0, 4.1, 8.0)

	if value != expected {
		t.Errorf("got %v, expected %v", value, expected)
	}
}
