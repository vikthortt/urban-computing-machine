package conversions

import (
	"testing"
)

func TestLoadLengthConversions(t *testing.T) {
	err := LoadLengthConversions("../data/length_conversions.json")
	if err != nil {
		t.Fatalf("Failed to load length conversions: %v", err)
	}
}

func TestLoadLengthConversionsInitializeLengthConversions(t *testing.T) {
	err := LoadLengthConversions("../data/length_conversions.json")
	if err != nil {
		t.Fatalf("Failed to load length conversions: %v", err)
	}

	// Assuming you have a function to check if LengthConversions is initialized
	if len(LengthConversions) == 0 {
		t.Fatal("LengthConversions is not initialized")
	}
}
