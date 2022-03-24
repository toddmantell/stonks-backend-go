package stonksbackend

import (
	"testing"
)

func TestGrahamFormula(t *testing.T) {
	expected := result{highGrahamFormulaNumber: 140.29594272076372, highConservativeGrahamFormulaNumber: 93.25059665871122, lowGrahamFormulaNumber: 323.436754176611, lowConservativeGrahamFormulaNumber: 184.82100238663486}
	value := GrahamFormula(15.0, 4.1, 8.0)

	t.Log("Running Test")

	if value != expected {
		t.Errorf("got %v, expected %v", value, expected)
		return
	}

	t.Log("PASSED")
}
