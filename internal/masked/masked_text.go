package masked

import (
	"log/slog"
	"strconv"

	"gopkg.in/yaml.v3"
)

type masker interface {
	DefaultMaskFunc() func(string) string
}

type MaskedText struct {
	s    string
	mask func(string) string
}

func maskOrFixed(s string, mask func(string) string) string {
	if mask == nil {
		return FixedMaskFunc(s)
	}
	return mask(s)
}

func NewMaskedText(s string, mask func(string) string) MaskedText {
	return MaskedText{s: s, mask: mask}
}

func (m MaskedText) Plain() string {
	return m.s
}

func (m MaskedText) String() string {
	return maskOrFixed(m.s, m.mask)
}

func (m MaskedText) LogValue() slog.Value {
	return slog.StringValue(maskOrFixed(m.s, m.mask))
}

func (m MaskedText) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(m.mask(m.s))), nil
}

func (m *MaskedText) UnmarshalJSON(v []byte) error {
	s, err := strconv.Unquote(string(v))

	if m.mask == nil {
		m.mask = FixedMaskFunc
	}
	m.s = s

	return err
}

func (m MaskedText) MarshalYAML() (interface{}, error) {
	return m.mask(m.s), nil
}

func (m *MaskedText) UnmarshalYAML(value *yaml.Node) error {
	var s string
	err := value.Decode(&s)

	if m.mask == nil {
		m.mask = FixedMaskFunc
	}
	m.s = s

	return err
}
