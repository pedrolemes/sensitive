package sensitive_test

import (
	"fmt"

	"github.com/pedrolemes/sensitive"
)

func Example() {
	// All sensitive types could use a custom mask function
	mask := func(s string) string {
		return "MASKED"
	}

	_ = sensitive.NewEmailWithMask("foo@bar.com", mask)
	_ = sensitive.NewPhoneWithMask("1234567890", mask)
	pass := sensitive.NewPasswordWithMask("my-password", mask)

	fmt.Println(pass)
	// Output: MASKED
}

func ExamplePassword() {
	p := sensitive.NewPassword("my-password")
	fmt.Println(p)
	// Output: ****
}
