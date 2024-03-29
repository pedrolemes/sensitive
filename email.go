package sensitive

import (
	"github.com/pedrolemes/sensitive/internal/masked"
)

// Email represents a sensitive email address.
type Email struct {
	masked.MaskedText
}

// NewEmail creates a new Email instance with the default email masking function.
func NewEmail(s string) Email {
	return NewEmailWithMask(s, masked.MaskEmailFunc)
}

// NewEmailWithMask creates a new Email instance with a custom masking function.
func NewEmailWithMask(s string, mask MaskFunc) Email {
	return Email{masked.NewMaskedText(s, mask)}
}
