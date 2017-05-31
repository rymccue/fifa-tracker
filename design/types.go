package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var LoginPayload = Type("LoginPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("password", String, func() {
		MinLength(5)
		MaxLength(100)
		Example("abcd1234")
	})
	Required("email", "password")
})

var UpdateMatchPayload = Type("UpdateMatchPayload", func() {
	Attribute("home_goals", Integer, func() {
		Minimum(0)
	})
	Attribute("away_goals", Integer, func() {
		Minimum(0)
	})
	Attribute("home_team", String, func() {
		MaxLength(200)
	})
	Attribute("away_team", String, func() {
		MaxLength(200)
	})
})

var MatchPayload = Type("MatchPayload", func() {
	Attribute("home_team", String, func() {
		MaxLength(200)
	})
	Attribute("away_team", String, func() {
		MaxLength(200)
	})
	Attribute("home_goals", Integer, func() {
		Minimum(0)
	})
	Attribute("away_goals", Integer, func() {
		Minimum(0)
	})
	Attribute("home_user", Integer, func() {
		Minimum(1)
	})
	Attribute("away_user", Integer, func() {
		Minimum(1)
	})

	Required("home_team", "away_team", "home_goals", "away_goals", "home_user", "away_user")
})

var AccountPayload = Type("AccountPayload", func() {
	Attribute("first_name", String, func() {
		MinLength(2)
		MaxLength(200)
		Example("James")
	})

	Attribute("last_name", String, func() {
		MinLength(2)
		MaxLength(200)
		Example("Bond")
	})

	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("Bond")
	})

	Attribute("password", String, func() {
		MinLength(5)
		MaxLength(100)
		Example("abcd1234")
	})
	Required("first_name", "last_name", "email", "password")
})

var UpdateAccountPayload = Type("UpdateAccountPayload", func() {
	Attribute("first_name", String, func() {
		MinLength(2)
		MaxLength(200)
		Example("James")
	})

	Attribute("last_name", String, func() {
		MinLength(2)
		MaxLength(200)
		Example("Bond")
	})

	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("Bond")
	})
})
