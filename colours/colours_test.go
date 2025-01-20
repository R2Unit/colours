package colours

import (
	"testing"
)

func TestColoursConstants(t *testing.T) {
	if Green == "" {
		t.Errorf("Expected Green to have a value, got empty string")
	}
	if Yellow == "" {
		t.Errorf("Expected Yellow to have a value, got empty string")
	}
	if Bold == "" {
		t.Errorf("Expected Bold to have a value, got empty string")
	}
	if Reset == "" {
		t.Errorf("Expected Reset to have a value, got empty string")
	}
}
