package services

import "testing"

// Unit Tests
func TestAdd(t *testing.T) {
	given := AddOperation(3, 5)
	expected := 8

	intNotEqual(given, expected, t)
}

func TestSub(t *testing.T) {
	given := SubOperation(5, 2)
	expected := 3

	intNotEqual(given, expected, t)

}

func intNotEqual(one int, two int, t *testing.T) {
	if one != two {
		t.Errorf("got %q, wanted %q", one, two)
	}
}

// Benchmarks
var intA = 300
var intB = 200

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddOperation(intA, intB)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubOperation(intA, intB)
	}
}
