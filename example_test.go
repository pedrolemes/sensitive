package sensitive_test

import (
	"fmt"

	"github.com/pedrolemes/sensitive"
)

func ExamplePassword() {
	p := sensitive.NewPassword("my-password")
	fmt.Println(p)
	// Output: ****
}

func ExampleNewPasswordWithMask() {
	p := sensitive.NewPasswordWithMask("my-password", func(s string) string {
		return "---"
	})

	fmt.Println(p)
	// Output: ---
}
