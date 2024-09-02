package sensitive_test

import (
	"encoding/json"
	"fmt"

	"github.com/pedrolemes/sensitive"
	"gopkg.in/yaml.v3"
)

func ExampleEmail() {
	e := sensitive.NewEmail("foo@bar.com")
	fmt.Println(e)
	// Output: f**@b**.com
}

func ExampleEmail_UnmarshalJSON() {
	type s struct {
		Email sensitive.Email `json:"email"`
	}

	data := []byte(`{"email":"foo@bar.com"}`)

	var v s
	_ = json.Unmarshal(data, &v)

	fmt.Println(v.Email)
	// Output: f**@b**.com
}

func ExampleEmail_UnmarshalYAML() {
	type s struct {
		Email sensitive.Email `yaml:"email"`
	}

	data := []byte(`email: foo@bar.com`)

	var v s
	_ = yaml.Unmarshal(data, &v)

	fmt.Println(v.Email)
	// Output: f**@b**.com
}
