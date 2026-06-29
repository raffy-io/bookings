package forms

import (
	"time"
)

type AvailableForm struct {
	Arrival   string
	Departure string
}

func ValidateAvailable(form *AvailableForm) (FormErrors, bool) {
	errors := FormErrors{}
	const dateLayout = "2006-01-02"
	var arrival, departure time.Time
	var err error

	// Normalize 'now' to UTC midnight for an accurate daily check
	today := time.Now().UTC().Truncate(24 * time.Hour)

	// Arrival Validation
	if form.Arrival == "" {
		errors["arrival"] = "Arrival date is required"
	} else {
		arrival, err = time.Parse(dateLayout, form.Arrival)
		if err != nil {
			errors["arrival"] = "Invalid arrival date format"
		} else if arrival.Before(today) {
			errors["arrival"] = "Arrival date cannot be in the past"
		}
	}

	// Departure Validation
	if form.Departure == "" {
		errors["departure"] = "Departure date is required"
	} else {
		departure, err = time.Parse(dateLayout, form.Departure)
		if err != nil {
			errors["departure"] = "Invalid departure date format"
		}
	}

	// Business Logic Cross-Field Validation
	if errors["arrival"] == "" && errors["departure"] == "" {
		if !departure.After(arrival) {
			errors["departure"] = "Departure date must be after the arrival date"
		}
	}

	return errors, errors.Valid()
}