package main

import (
	"encoding/json"

	"time"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/nachkonvention/go-schedule/model"
)

var appointments []model.Appointment = []model.Appointment{
	model.NewAppointment("foo", "bar", time.Now(), time.Now(), model.Talk),
	model.NewAppointment("bar", "foo", time.Now(), time.Now(), model.Talk),
}

func main() {
	m := martini.Classic()

	m.Get("/api/appointments", func() string {
		appointment, _ := json.Marshal(appointments)
		return string(appointment)
	})

	m.Post("/api/appointments", binding.Bind(model.Appointment{}), func(appointment model.Appointment) string {
		appointments = append(appointments, appointment)
		result, _ := json.Marshal(appointment)
		return string(result)
	})

	m.Run()
}
