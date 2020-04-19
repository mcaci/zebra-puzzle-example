package zebra

import "testing"

func TestSolvePuzzle(t *testing.T) {
	expected := Solution{DrinksWater: "Norwegian", OwnsZebra: "Japanese"}
	actual := SolvePuzzle()
	if expected != actual {
		t.Fatalf("FAILED:\nExpected: %#v\nActual: %#v",
			expected, actual)
	}
}
