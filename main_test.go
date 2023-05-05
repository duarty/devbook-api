package main

import "testing"

func TestSum(t *testing.T) {

	type ValuesToSum struct {
		a int
		b int
	}

	values := []ValuesToSum{
		{
			a: 2,
			b: 2,
		},
	}

	for _, value := range values {
		result := Sum(value.a, value.b)
		if result != 4 {
			t.Error("FAIL")
		}
	}
}
