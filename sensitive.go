// package sensitive provides a way to create sensitive data types
package sensitive

// MaskFunc is a function type that takes a string as input and returns a masked string.
// It is used to define custom masking functions for sensitive data.
type MaskFunc func(string) string
