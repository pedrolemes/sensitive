package sensitive

import "github.com/pedrolemes/sensitive/internal/masked"

// Phone represents a sensitive phone number.
type Phone struct {
	masked.MaskedText
}

// NewPhone creates a new Phone instance with the default mask.
func NewPhone(s string) Phone {
	return NewPhoneWithMask(s, masked.MaskPhoneFunc)
}

// NewPhoneWithMask creates a new Phone instance with a custom mask.
func NewPhoneWithMask(s string, mask MaskFunc) Phone {
	return Phone{masked.NewMaskedText(s, mask)}
}
