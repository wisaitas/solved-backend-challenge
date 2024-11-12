package three

import (
	"encoding/json"
	"testing"
)

func TestPireFireDire(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "basic word count",
			input: "pork beef pork chicken beef",
			expected: map[string]int{
				"pork":    2,
				"beef":    2,
				"chicken": 1,
			},
		},
		{
			name:  "with punctuation",
			input: "beef, pork. beef",
			expected: map[string]int{
				"beef": 2,
				"pork": 1,
			},
		},
		{
			name:  "case insensitive",
			input: "Beef PORK beef",
			expected: map[string]int{
				"beef": 2,
				"pork": 1,
			},
		},
		{
			name:     "empty input",
			input:    "",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PireFireDire(tt.input)

			if tt.input == "" {
				var response map[string]interface{}
				err := json.Unmarshal([]byte(result), &response)
				if err != nil {
					t.Errorf("Failed to unmarshal result: %v", err)
				}
				return
			}

			var response map[string]interface{}
			err := json.Unmarshal([]byte(result), &response)
			if err != nil {
				t.Errorf("Failed to unmarshal result: %v", err)
			}

			beefCount, ok := response["beef"].(map[string]interface{})
			if !ok {
				t.Error("Expected beef count map in response")
			}

			for word, expectedCount := range tt.expected {
				count, exists := beefCount[word]
				if !exists {
					t.Errorf("Expected word %q not found in result", word)
					continue
				}
				if float64(expectedCount) != count {
					t.Errorf("For word %q: expected count %d, got %v", word, expectedCount, count)
				}
			}
		})
	}
}
