package colours

import "testing"

func TestColoursConstants(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"Green", Green},
		{"Yellow", Yellow},
		{"Red", Red},
		{"Blue", Blue},
		{"Cyan", Cyan},
		{"Magenta", Magenta},
		{"Bold", Bold},
		{"Reset", Reset},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value == "" {
				t.Errorf("Expected %s to have a value, but got an empty string", tt.name)
			}
		})
	}
}
