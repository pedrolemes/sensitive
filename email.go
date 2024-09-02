package sensitive

import (
	"github.com/pedrolemes/sensitive/internal/masked"
	"gopkg.in/yaml.v3"
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

func (e *Email) UnmarshalJSON(v []byte) error {
	e.MaskedText = masked.NewMaskedText("", masked.MaskEmailFunc)
	return e.MaskedText.UnmarshalJSON(v)
}

func (e *Email) UnmarshalYAML(value *yaml.Node) error {
	e.MaskedText = masked.NewMaskedText("", masked.MaskEmailFunc)
	return e.MaskedText.UnmarshalYAML(value)
}
