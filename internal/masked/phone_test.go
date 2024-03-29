package masked_test

import (
	"testing"

	"github.com/pedrolemes/sensitive/internal/masked"
)

func TestMaskPhoneFunc(t *testing.T) {
	type testData struct {
		input    string
		expected string
	}

	tests := []testData{
		{input: "11999999999", expected: "1199999****"},
		{input: "1199999999", expected: "119999****"},
		{input: "119999999", expected: "11999****"},
		{input: "11999999", expected: "1199****"},
		{input: "1199999", expected: "119****"},
		{input: "119999", expected: "11****"},
		{input: "11999", expected: "1****"},
		{input: "1199", expected: "****"},
		{input: "119", expected: "****"},
		{input: "11", expected: "****"},
		{input: "1", expected: "****"},
		{input: "", expected: "****"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := masked.MaskPhoneFunc(test.input)
			if actual != test.expected {
				t.Errorf("expected %s, got %s", test.expected, actual)
			}
		})
	}
}
