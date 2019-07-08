package main

import (
	"math"
)

func SquareByMultiplying(x float64) float64 {
	return x * x
}

func SquareWithPow(x float64) float64 {
	return math.Pow(x, 2)
}

func QuadByMultiplying(x float64) float64 {
	return x * x * x * x
}

func QuadWithPow(x float64) float64 {
	return math.Pow(x, 4)
}
