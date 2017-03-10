package main

import (
	"encoding/json"
	"log"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/nachkonvention/go-schedule/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	m := martini.Classic()

	m.Get("/api/appointments", func() string {
		session := getDatabaseSession()
		defer session.Close()

		var appointments []model.Appointment
		session.DB("go-schedule").C("appointments").Find(bson.M{}).All(&appointments)

		appointment, _ := json.Marshal(appointments)
		return string(appointment)
	})

	m.Post("/api/appointments", binding.Bind(model.Appointment{}), func(appointment model.Appointment) string {
		session := getDatabaseSession()
		defer session.Close()

		appointmentsCollection := session.DB("go-schedule").C("appointments")

		err := appointmentsCollection.Insert(appointment)
		if err != nil {
			log.Fatal(err)
		}

		result, _ := json.Marshal(appointment)
		return string(result)
	})

	m.Run()
}

func getDatabaseSession() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	return session
}
