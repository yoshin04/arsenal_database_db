package utils

func OptionalString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
