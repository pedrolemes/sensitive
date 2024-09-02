package sensitive_test

import (
	"encoding/json"
	"fmt"

	"github.com/pedrolemes/sensitive"
	"gopkg.in/yaml.v3"
)

func ExamplePhone() {
	p := sensitive.NewPhone("1234567890")
	fmt.Println(p)
	// Output: 123456****
}

func ExamplePhone_UnmarshalJSON() {
	type s struct {
		Phone sensitive.Phone `json:"phone"`
	}

	data := []byte(`{"phone":"1234567890"}`)

	var v s
	_ = json.Unmarshal(data, &v)

	fmt.Println(v.Phone)
	// Output: 123456****
}

func ExamplePhone_UnmarshalYAML() {
	type s struct {
		Phone sensitive.Phone `yaml:"phone"`
	}

	data := []byte(`phone: 1234567890`)

	var v s
	_ = yaml.Unmarshal(data, &v)

	fmt.Println(v.Phone)
	// Output: 123456****
}
