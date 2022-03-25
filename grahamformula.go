package main

type Result struct {
	highGrahamFormulaNumber             float64
	highConservativeGrahamFormulaNumber float64
	lowGrahamFormulaNumber              float64
	lowConservativeGrahamFormulaNumber  float64
}

const REGULAR_PE = 8.5
const CONSERVATIVE_PE = 7
const REGULAR_GROWTH_MULTIPLE = 2
const CONSERVATIVE_GROWTH_MULTIPLE = 1

func GrahamFormula(lowGrowthRate float64, highGrowthRate float64, ttmEPS float64) Result {
	high := calcGrahamFormula(highGrowthRate, ttmEPS, REGULAR_PE, REGULAR_GROWTH_MULTIPLE)
	highConservative := calcGrahamFormula(highGrowthRate, ttmEPS, CONSERVATIVE_PE, CONSERVATIVE_GROWTH_MULTIPLE)
	low := calcGrahamFormula(lowGrowthRate, ttmEPS, REGULAR_PE, REGULAR_GROWTH_MULTIPLE)
	lowConservative := calcGrahamFormula(lowGrowthRate, ttmEPS, CONSERVATIVE_PE, CONSERVATIVE_GROWTH_MULTIPLE)
	r := Result{
		highGrahamFormulaNumber:             high,
		highConservativeGrahamFormulaNumber: highConservative,
		lowGrahamFormulaNumber:              low,
		lowConservativeGrahamFormulaNumber:  lowConservative,
	}

	return r
}

func calcGrahamFormula(growthRate, ttmEPS, peAssumption, growthMultiple float64) float64 {
	const DISCOUNTRATE = 4.4
	const CURRENTAAARATE = 1.86
	const LONGTERMAAARATE = 4.19

	value := (ttmEPS * (peAssumption + growthMultiple*growthRate) * DISCOUNTRATE) / LONGTERMAAARATE
	return value
}
