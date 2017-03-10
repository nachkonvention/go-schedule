package model

import (
	"time"
)

// Appointment represents an appointment
type Appointment struct {
	Title           string
	Starttime       time.Time
	Endtime         time.Time
	Description     string
	AppointmentType AppointmentType
}

// NewAppointment creates a new appointment
func NewAppointment(title, description string, starttime, endtime time.Time, appointmentType AppointmentType) (appointment Appointment) {
	appointment = Appointment{Title: title, Description: description, Starttime: starttime, Endtime: endtime, AppointmentType: appointmentType}
	return
}
