package masked_test

import (
	"testing"

	"github.com/pedrolemes/sensitive/internal/masked"
)

const handled = "76e45cfd-268f-44d0-9af9-a5515366870f"
const sensitive = "f2fb018f-5505-4c33-b72c-269dcf678de0"

func testMaskFunc(string) string {
	return handled
}

func TestPlain(t *testing.T) {
	t.Run("should return the plain text", func(t *testing.T) {
		m := masked.NewMaskedText(sensitive, testMaskFunc)

		result := m.Plain()

		if result != sensitive {
			t.Errorf("expected %q, got %q", sensitive, result)
		}
	})
}

func TestString(t *testing.T) {
	t.Run("should return the masked text", func(t *testing.T) {
		m := masked.NewMaskedText(sensitive, testMaskFunc)

		result := m.String()

		if result != handled {
			t.Errorf("expected %q, got %q", handled, result)
		}
	})
}
