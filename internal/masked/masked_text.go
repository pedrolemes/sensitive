package masked

import (
	"log/slog"
	"strconv"

	"gopkg.in/yaml.v3"
)

type MaskedText struct {
	s    string
	mask func(string) string
}

func NewMaskedText(s string, mask func(string) string) MaskedText {
	return MaskedText{s: s, mask: mask}
}

func (m MaskedText) Plain() string {
	return m.s
}

func (m MaskedText) String() string {
	return m.mask(m.s)
}

func (m MaskedText) LogValue() slog.Value {
	return slog.StringValue(m.mask(m.s))
}

func (m MaskedText) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(m.mask(m.s))), nil
}

func (m *MaskedText) UnmarshalJSON(v []byte) error {
	s, err := strconv.Unquote(string(v))
	*m = MaskedText{
		s:    s,
		mask: FixedMaskFunc,
	}

	return err
}

func (m MaskedText) MarshalYAML() (interface{}, error) {
	return m.mask(m.s), nil
}

func (m *MaskedText) UnmarshalYAML(value *yaml.Node) error {
	var s string
	err := value.Decode(&s)

	*m = MaskedText{
		s:    s,
		mask: FixedMaskFunc,
	}

	return err
}
