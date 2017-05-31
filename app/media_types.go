// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "FIFA Tracker": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/rymccue/fifa-tracker/design
// --out=$(GOPATH)/src/github.com/rymccue/fifa-tracker
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// A user account (default view)
//
// Identifier: application/ft.account+json; view=default
type FtAccount struct {
	// Email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// First name of the user
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	// ID of the user
	ID int `form:"id" json:"id" xml:"id"`
	// Last name of the user
	LastName string `form:"last_name" json:"last_name" xml:"last_name"`
}

// Validate validates the FtAccount media type instance.
func (mt *FtAccount) Validate() (err error) {

	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if mt.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	return
}

// FtAccountCollection is the media type for an array of FtAccount (default view)
//
// Identifier: application/ft.account+json; type=collection; view=default
type FtAccountCollection []*FtAccount

// Validate validates the FtAccountCollection media type instance.
func (mt FtAccountCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A FIFA Match (default view)
//
// Identifier: application/ft.match+json; view=default
type FtMatch struct {
	// The away goals
	AwayGoals int `form:"away_goals" json:"away_goals" xml:"away_goals"`
	// The away team
	AwayTeam string `form:"away_team" json:"away_team" xml:"away_team"`
	// The away user
	AwayUser int `form:"away_user" json:"away_user" xml:"away_user"`
	// The home goals
	HomeGoals int `form:"home_goals" json:"home_goals" xml:"home_goals"`
	// The home team
	HomeTeam string `form:"home_team" json:"home_team" xml:"home_team"`
	// The home user
	HomeUser int `form:"home_user" json:"home_user" xml:"home_user"`
}

// Validate validates the FtMatch media type instance.
func (mt *FtMatch) Validate() (err error) {
	if mt.HomeTeam == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "home_team"))
	}
	if mt.AwayTeam == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "away_team"))
	}

	return
}

// FtMatchCollection is the media type for an array of FtMatch (default view)
//
// Identifier: application/ft.match+json; type=collection; view=default
type FtMatchCollection []*FtMatch

// Validate validates the FtMatchCollection media type instance.
func (mt FtMatchCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A JWT for a user (default view)
//
// Identifier: application/ft.token+json; view=default
type FtToken struct {
	// A JWT token
	Token string `form:"token" json:"token" xml:"token"`
}

// Validate validates the FtToken media type instance.
func (mt *FtToken) Validate() (err error) {
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	return
}
