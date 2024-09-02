package masked_test

import (
	"bytes"
	"testing"

	"github.com/pedrolemes/sensitive/internal/masked"
	"github.com/spf13/viper"
)

func TestViper(t *testing.T) {
	var input = []byte(`m: some sensitive data`)

	viper.SetConfigType("yaml")
	_ = viper.ReadConfig(bytes.NewBuffer(input))

	var s someStruct
	err := viper.Unmarshal(&s, viper.DecodeHook(masked.MapStructureDecodeHook()))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if s.M.String() != "****" {
		t.Errorf("expected %q, got %q", "****", s.M.String())
	}

	if s.M.Plain() != "some sensitive data" {
		t.Errorf("expected %q, got %q", "some sensitive data", s.M.Plain())
	}
}
