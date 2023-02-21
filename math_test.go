package main

import "testing"

func TestSoma(t *testing.T) {
	expected := 30
	result := Soma(15, 15)

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
