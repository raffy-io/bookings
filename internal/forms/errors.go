package forms

// FormErrors holds validation errors for specific fields across any form
type FormErrors map[string]string

// Valid returns true if there are no errors
func (e FormErrors) Valid() bool {
	return len(e) == 0
}
