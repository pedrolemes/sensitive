package masked

// MaskPhoneFunc masks the last four digits of a phone number.
// If the input string is less than four characters, it returns "****".
// Otherwise, it replaces the last four characters with "****".
func MaskPhoneFunc(s string) string {
	if len(s) < 4 {
		return fixedMask
	}

	return s[:len(s)-4] + "****"
}
