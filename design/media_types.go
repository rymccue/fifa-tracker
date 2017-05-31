package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Account is the account resource media type.
var Account = MediaType("application/ft.account+json", func() {
	Description("A user account")
	Attributes(func() {
		Attribute("id", Integer, "ID of the user", func() {
			Example(1)
		})
		Attribute("first_name", String, "First name of the user", func() {
			Example("James")
		})
		Attribute("last_name", String, "Last name of the user", func() {
			Example("Bond")
		})
		Attribute("email", String, "Email of the user", func() {
			Example("james@bond.com")
		})
		Required("id", "first_name", "last_name", "email")
	})

	View("default", func() {
		Attribute("id")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("email")
	})
})

var Match = MediaType("application/ft.match+json", func() {
	Description("A FIFA Match")
	Attributes(func() {
		Attribute("home_team", String, "The home team", func() {
			Example("Real Madrid")
		})
		Attribute("away_team", String, "The away team", func() {
			Example("Barceloca")
		})
		Attribute("home_user", Integer, "The home user", func() {
			Example(1)
		})
		Attribute("away_user", Integer, "The away user", func() {
			Example(2)
		})
		Attribute("away_goals", Integer, "The away goals", func() {
			Example(2)
		})
		Attribute("home_goals", Integer, "The home goals", func() {
			Example(1)
		})
		Required("home_team", "away_team", "home_user", "away_user", "home_goals", "away_goals")
	})

	View("default", func() {
		Attribute("home_team")
		Attribute("away_team")
		Attribute("home_user")
		Attribute("away_user")
		Attribute("home_goals")
		Attribute("away_goals")
	})
})

var Token = MediaType("application/ft.token+json", func() {
	Description("A JWT for a user")
	Attributes(func() {
		Attribute("token", String, "A JWT token")
		Required("token")
	})

	View("default", func() {
		Attribute("token")
	})
})
