package messages

import (
	"strings"
	"testing"

	"github.com/r2unit/eyes/colours"
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
		{"Danger", Danger, "[Danger]", "An error occurred"},
		{"Info", Info, "[Info]", "Informational message"},
		{"Debug", Debug, "[Debug]", "Debugging details"},
		{"Notice", Notice, "[Notice]", "Important notice"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.function(tc.message)

			// Check that the result contains the prefix
			if !strings.Contains(result, tc.prefix) {
				t.Errorf("Expected result to contain prefix %s, got %s", tc.prefix, result)
			}

			// Check that the result contains the message
			if !strings.Contains(result, tc.message) {
				t.Errorf("Expected result to contain message %s, got %s", tc.message, result)
			}

			// Check that the result starts with a valid timestamp
			if !strings.HasPrefix(result, getTimestamp()[:10]) { // Only validate the date part
				t.Errorf("Expected result to start with a valid timestamp, got %s", result)
			}

			// Check that the result ends with the Reset code
			if !strings.HasSuffix(result, colours.Reset) {
				t.Errorf("Expected result to end with Reset, got %s", result)
			}
		})
	}
}
