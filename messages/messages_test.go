package messages

import (
	"strings"
	"testing"

	"github.com/r2unit/go-colours/colours"
)

func TestColourize(t *testing.T) {
	message := "Hello, World!"
	colouredMessage := Colourize(colours.Green, message)

	if !strings.Contains(colouredMessage, message) {
		t.Errorf("Expected coloured message to contain the input message, got %s", colouredMessage)
	}

	if !strings.HasPrefix(colouredMessage, colours.Green) {
		t.Errorf("Expected coloured message to start with the color code, got %s", colouredMessage)
	}

	if !strings.HasSuffix(colouredMessage, colours.Reset) {
		t.Errorf("Expected coloured message to end with Reset, got %s", colouredMessage)
	}
}

func TestPredefinedMessages(t *testing.T) {
	testCases := []struct {
		name     string
		function func(string) string
		prefix   string
		message  string
	}{
		{"Success", Success, "[Success]", "Operation successful"},
		{"Warning", Warning, "[Warning]", "Potential issue"},
		{"Error", Error, "[Error]", "An error occurred"},
		{"Info", Info, "[Info]", "Informational message"},
		{"Debug", Debug, "[Debug]", "Debugging details"},
		{"Notice", Notice, "[Notice]", "Important notice"},
		{"Ok", Ok, "[Ok]", "Ok Operation"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.function(tc.message)

			if !strings.Contains(result, tc.prefix) {
				t.Errorf("Expected result to contain prefix %s, got %s", tc.prefix, result)
			}

			if !strings.Contains(result, tc.message) {
				t.Errorf("Expected result to contain message %s, got %s", tc.message, result)
			}

			if !strings.HasPrefix(result, colours.Bold) {
				t.Errorf("Expected result to start with Bold, got %s", result)
			}

			if !strings.HasSuffix(result, colours.Reset) {
				t.Errorf("Expected result to end with Reset, got %s", result)
			}
		})
	}
}
