package models

import "github.com/raffy-io/bookings/internal/forms"

type NotifType struct {
	SuccessMessage string
	ErrorMessage   string
	ErrorsMap  forms.FormErrors
}

type ReservationSummary struct {
	Name string
	Arrival string
	Departure string
	Email string
	Phone string
}