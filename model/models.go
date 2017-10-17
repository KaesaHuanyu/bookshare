package model

import "gopkg.in/mgo.v2/bson"

type (
	Location struct {
		Name string
		Address string
		Latitude float64
		Longitude float64
	}

	Book struct {
		Id bson.ObjectId
		Name string
		ISBN string
	}

	User struct {
		Location *Location
		Id bson.ObjectId

	}
)
