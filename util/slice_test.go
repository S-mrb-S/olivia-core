package util

import (
	"testing"
)

func TestContains(t *testing.T) {
	slice := []string{"hey", "Hola", "boNjour"}

	if !SliceIncludes(slice, "Hola") || SliceIncludes(slice, "bonjour") || !SliceIncludes(slice, "hey") {
		t.Errorf("slice.Contains() failed.")
	}
}

func TestDifference(t *testing.T) {
	slice := []string{"hey", "Hola", "boNjour"}
	slice2 := []string{"hey", "boNjour"}
	difference := SliceDifference(slice, slice2)

	if len(difference) != 1 || difference[0] != "Hola" {
		t.Errorf("Difference() failed.")
	}
}

func TestIndex(t *testing.T) {
	slice := []string{"hey", "Hola", "boNjour"}

	if SliceIndex(slice, "Hola") != 1 || SliceIndex(slice, "boNjour") != 2 {
		t.Errorf("slice.Index() failed.")
	}
}
