package integers

import (
	"testing"
)

func TestAdders(t *testing.T) {
	result := Add(10, 20)
	expected := 30

	assertResultMessage(t, result, expected)
}

func TestSubsctract(t *testing.T) {
	result := Substract(20, 10)
	expected := 10

	assertResultMessage(t, result, expected)
}

func TestDivide(t *testing.T) {
	result := Divide(20, 10)
	expected := 2

	assertResultMessage(t, result, expected)
}

func TestMultiply(t *testing.T) {
	result := Multiply(20, 10)
	expected := 200

	assertResultMessage(t, result, expected)
}

func assertResultMessage(t testing.TB, result, expected int) {
	t.Helper()

	if result != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, result)
	}
}
