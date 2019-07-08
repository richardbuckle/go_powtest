package main

import "testing"

func TestSquareByMultiplying(t *testing.T) {
	x := 5.0
	y := SquareByMultiplying(x)
	if y != 25.0 {
		t.Errorf(`SquareByMultiplying(%v) = %v, expected 25.0`, x, y)
	}
}

func TestSquareWithPow(t *testing.T) {
	x := 5.0
	y := SquareWithPow(x)
	if y != 25.0 {
		t.Errorf(`SquareWithPow(%v) = %v, expected 25.0`, x, y)
	}
}

func BenchmarkSquareByMultiplying(b *testing.B) {
	x := 5.0
	for i := 0; i < b.N; i++ {
		_ = SquareByMultiplying(x)
	}
}

func BenchmarkSquareWithPow(b *testing.B) {
	x := 5.0
	for i := 0; i < b.N; i++ {
		_ = SquareWithPow(x)
	}
}

func BenchmarkQuadByMultiplying(b *testing.B) {
	x := 5.0
	for i := 0; i < b.N; i++ {
		_ = QuadByMultiplying(x)
	}
}

func BenchmarkQuadWithPow(b *testing.B) {
	x := 5.0
	for i := 0; i < b.N; i++ {
		_ = QuadWithPow(x)
	}
}
