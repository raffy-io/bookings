package forms

import (
	"net/mail"
	"regexp"
	"unicode/utf8"
)

// Compile the regex ONCE at the package level so it's highly performant.
// This pattern matches basic phone numbers: e.g., +1234567890, 0912-345-6789, 123 456 7890
var phoneRegex = regexp.MustCompile(`^[+]?[0-9\s\-]{7,15}$`)

// ReserveForm  will hold data
type ReserveForm struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

func ValidateReservation(form *ReserveForm) (FormErrors, bool) {
	errors := FormErrors{}

	// First name Validation
	if form.FirstName == "" {
		errors["first_name"] = "First name is required"
	} else if utf8.RuneCountInString(form.FirstName) < 2 {
		errors["first_name"] = "First name must be at least 2 characters long"
	}

	// Last name Validation
	if form.LastName == "" {
		errors["last_name"] = "Last name is required"
	} else if utf8.RuneCountInString(form.LastName) < 2 {
		errors["last_name"] = "Last name must be at least 2 characters long"
	}

	// Email Validation
	if form.Email == "" {
		errors["email"] = "Email is required"
	} else if _, err := mail.ParseAddress(form.Email); err != nil {
		errors["email"] = "Please enter a valid email address"
	}

	// Phone Validation
	if form.Phone == "" {
		errors["phone"] = "Phone number is required"
	} else if !phoneRegex.MatchString(form.Phone) {
		errors["phone"] = "Please enter a valid phone number (numbers, spaces, or dashes only)"
	}

	return errors, errors.Valid()
}
