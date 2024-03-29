package sensitive

import "github.com/pedrolemes/sensitive/internal/masked"

// Password represents a sensitive password value.
type Password struct {
	masked.MaskedText
}

// NewPassword creates a new Password instance with a default mask.
func NewPassword(s string) Password {
	return NewPasswordWithMask(s, masked.FixedMaskFunc)
}

// NewPasswordWithMask creates a new Password instance with a custom mask.
func NewPasswordWithMask(s string, mask MaskFunc) Password {
	return Password{masked.NewMaskedText(s, mask)}
}
