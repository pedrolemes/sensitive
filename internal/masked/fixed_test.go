package masked_test

import (
	"testing"

	"github.com/pedrolemes/sensitive/internal/masked"
)

func TestFixedMaskFunc(t *testing.T) {
	args := []string{
		"123456",
		"",
		"83916744-cd19-4eb0-9863-bbc03b2eeb7a",
	}

	for _, arg := range args {
		t.Run(arg, func(t *testing.T) {
			actual := masked.FixedMaskFunc(arg)
			expected := "****"
			if actual != expected {
				t.Errorf("expected %s, got %s", expected, actual)
			}
		})
	}
}
