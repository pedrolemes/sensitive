package masked

import "strings"

// MaskEmailFunc masks the email address by replacing characters with asterisks.
// It takes a string parameter representing the email address and returns the masked email address.
func MaskEmailFunc(s string) string {
	at := strings.Index(s, "@")

	if at > 0 {
		lastDot := strings.LastIndex(s, ".")
		if lastDot == -1 {
			lastDot = len(s)
		}

		return s[:1] + strings.Repeat("*", at-1) + s[at:at+2] + strings.Repeat("*", lastDot-at-2) + s[lastDot:]
	}

	return fixedMask
}
