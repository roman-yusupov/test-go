package main

import (
	"math/rand"
	"testing"
)

func TestFactorial(t *testing.T) {

	s := &server{}

	t.Run("Factorial from negative numbers", func(t *testing.T) {

		f := s.factorial(-1)
		expected := "Number should be greater or equal to 0"
		if f != expected {
			t.Fatalf("Expected \"%s\", returned \"%s\"", expected, f)
		}

	})

	t.Run("Factorial from 0", func(t *testing.T) {

		f := s.factorial(0)
		expected := "0"
		if f != expected {
			t.Fatalf("Expected \"%s\", returned \"%s\"", expected, f)
		}

	})

	t.Run("Factorial from 1", func(t *testing.T) {

		f := s.factorial(1)
		expected := "1"
		if f != expected {
			t.Fatalf("Expected \"%s\", returned \"%s\"", expected, f)
		}

	})

	t.Run("Factorial from 5", func(t *testing.T) {

		f := s.factorial(5)
		expected := "120"
		if f != expected {
			t.Fatalf("Expected \"%s\", returned \"%s\"", expected, f)
		}

	})

	t.Run("Factorial from 99", func(t *testing.T) {

		f := s.factorial(99)
		expected := "933262154439441526816992388562667004907159682643816214685929638952175999932299156089414639761565182862536979208272237582511852109168640000000000000000000000"
		if f != expected {
			t.Fatalf("Expected \"%s\", returned \"%s\"", expected, f)
		}

	})

	t.Run("Factorial from 100", func(t *testing.T) {

		f := s.factorial(100)
		expected := "9.33262154439441e+157"
		if f != expected {
			t.Fatalf("Expected \"%s\", returned \"%s\"", expected, f)
		}

	})

	t.Run("Factorial from 1000", func(t *testing.T) {

		f := s.factorial(1000)
		expected := "+Inf"
		if f != expected {
			t.Fatalf("Expected \"%s\", returned \"%s\"", expected, f)
		}

	})

	t.Run("Factorial from 10000000", func(t *testing.T) {

		f := s.factorial(10000000)
		expected := "+Inf"
		if f != expected {
			t.Fatalf("Expected \"%s\", returned \"%s\"", expected, f)
		}

	})

}

func BenchmarkFactorial(b *testing.B) {
	s := &server{}

	for i := 0; i < b.N; i++ {
		s.factorial(rand.Int63n(1000))
	}
}
