package masked_test

import (
	"encoding/json"
	"testing"

	"github.com/pedrolemes/sensitive/internal/masked"
	"gopkg.in/yaml.v3"
)

type someStruct struct {
	M masked.MaskedText
}

func TestJson(t *testing.T) {
	t.Run("should marshal the masked text", func(t *testing.T) {
		s := someStruct{
			M: masked.NewMaskedText("some sensitive data", masked.FixedMaskFunc),
		}

		jsonStr, err := json.Marshal(s)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		expected := `{"M":"****"}`
		if string(jsonStr) != expected {
			t.Errorf("expected %q, got %q", expected, string(jsonStr))
		}
	})

	t.Run("should return error when could not unmarshal from json", func(t *testing.T) {
		jsonStr := `{"M": 1}`

		var s someStruct
		err := json.Unmarshal([]byte(jsonStr), &s)

		if err == nil {
			t.Errorf("expected an error, got nil")
			t.FailNow()
		}

		if s.M.String() != "****" {
			t.Errorf("expected %q, got %q", "****", s.M.String())
		}

		if s.M.Plain() != "" {
			t.Errorf("expected %q, got %q", "", s.M.Plain())
		}
	})

	t.Run("should unmarshal from json", func(t *testing.T) {
		jsonStr := `{"M":"some sensitive data"}`

		var s someStruct
		err := json.Unmarshal([]byte(jsonStr), &s)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if s.M.String() != "****" {
			t.Errorf("expected %q, got %q", "****", s.M.String())
		}

		if s.M.Plain() != "some sensitive data" {
			t.Errorf("expected %q, got %q", "some sensitive data", s.M.Plain())
		}
	})
}

func TestYaml(t *testing.T) {
	t.Run("should marshal the masked text", func(t *testing.T) {
		s := someStruct{
			M: masked.NewMaskedText("some sensitive data", masked.FixedMaskFunc),
		}

		yamlStr, err := yaml.Marshal(s)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		expected := "m: '****'\n"
		if string(yamlStr) != expected {
			t.Errorf("expected %q, got %q", expected, string(yamlStr))
		}
	})

	t.Run("should return error when could not unmarshal from yaml", func(t *testing.T) {
		yamlStr := `
m:
  - sensitive data`

		var s someStruct
		err := yaml.Unmarshal([]byte(yamlStr), &s)

		if err == nil {
			t.Errorf("expected an error, got nil")
			t.FailNow()
		}

		if s.M.String() != "****" {
			t.Errorf("expected %q, got %q", "****", s.M.String())
		}

		if s.M.Plain() != "" {
			t.Errorf("expected %q, got %q", "", s.M.Plain())
		}
	})

	t.Run("should unmarshal from yaml", func(t *testing.T) {
		yamlStr := `m: sensitive data`

		var s someStruct
		err := yaml.Unmarshal([]byte(yamlStr), &s)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if s.M.String() != "****" {
			t.Errorf("expected %q, got %q", "****", s.M.String())
		}

		if s.M.Plain() != "sensitive data" {
			t.Errorf("expected %q, got %q", "sensitive data", s.M.Plain())
		}
	})
}
