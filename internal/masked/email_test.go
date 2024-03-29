package masked_test

import (
	"testing"

	"github.com/pedrolemes/sensitive/internal/masked"
)

func TestMaskEmailFunc(t *testing.T) {
	type testData struct {
		input    string
		expected string
	}

	tests := []testData{
		{input: "valid@email.com", expected: "v****@e****.com"},
		{input: "invalid@invalid", expected: "i******@i******"},
		{input: "invalid", expected: "****"},
		{input: "invalid@invalid@com", expected: "i******@i**********"},
		// {input: "a@b.com", expected: "****@****"},
		// {input: "a@.c", expected: "****@****"},
		// {input: "a.b@.c", expected: "a**@****"},
		// {input: "a.b@c.com", expected: "a**@*.com"},
		{input: "", expected: "****"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := masked.MaskEmailFunc(test.input)
			if actual != test.expected {
				t.Errorf("expected %s, got %s", test.expected, actual)
			}
		})
	}
}
