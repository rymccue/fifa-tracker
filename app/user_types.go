// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "FIFA Tracker": Application User Types
//
// Command:
// $ goagen
// --design=github.com/rymccue/fifa-tracker/design
// --out=$(GOPATH)/src/github.com/rymccue/fifa-tracker
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// accountPayload user type.
type accountPayload struct {
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName  *string `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
	Password  *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the accountPayload type instance.
func (ut *accountPayload) Validate() (err error) {
	if ut.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if ut.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 6, true))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 400, false))
		}
	}
	if ut.FirstName != nil {
		if utf8.RuneCountInString(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, utf8.RuneCountInString(*ut.FirstName), 2, true))
		}
	}
	if ut.FirstName != nil {
		if utf8.RuneCountInString(*ut.FirstName) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, utf8.RuneCountInString(*ut.FirstName), 200, false))
		}
	}
	if ut.LastName != nil {
		if utf8.RuneCountInString(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, utf8.RuneCountInString(*ut.LastName), 2, true))
		}
	}
	if ut.LastName != nil {
		if utf8.RuneCountInString(*ut.LastName) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, utf8.RuneCountInString(*ut.LastName), 200, false))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 5, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 100, false))
		}
	}
	return
}

// Publicize creates AccountPayload from accountPayload
func (ut *accountPayload) Publicize() *AccountPayload {
	var pub AccountPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.FirstName != nil {
		pub.FirstName = *ut.FirstName
	}
	if ut.LastName != nil {
		pub.LastName = *ut.LastName
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// AccountPayload user type.
type AccountPayload struct {
	Email     string `form:"email" json:"email" xml:"email"`
	FirstName string `form:"first_name" json:"first_name" xml:"first_name"`
	LastName  string `form:"last_name" json:"last_name" xml:"last_name"`
	Password  string `form:"password" json:"password" xml:"password"`
}

// Validate validates the AccountPayload type instance.
func (ut *AccountPayload) Validate() (err error) {
	if ut.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "first_name"))
	}
	if ut.LastName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "last_name"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, ut.Email, goa.FormatEmail, err2))
	}
	if utf8.RuneCountInString(ut.Email) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, ut.Email, utf8.RuneCountInString(ut.Email), 6, true))
	}
	if utf8.RuneCountInString(ut.Email) > 400 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, ut.Email, utf8.RuneCountInString(ut.Email), 400, false))
	}
	if utf8.RuneCountInString(ut.FirstName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, ut.FirstName, utf8.RuneCountInString(ut.FirstName), 2, true))
	}
	if utf8.RuneCountInString(ut.FirstName) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, ut.FirstName, utf8.RuneCountInString(ut.FirstName), 200, false))
	}
	if utf8.RuneCountInString(ut.LastName) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, ut.LastName, utf8.RuneCountInString(ut.LastName), 2, true))
	}
	if utf8.RuneCountInString(ut.LastName) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, ut.LastName, utf8.RuneCountInString(ut.LastName), 200, false))
	}
	if utf8.RuneCountInString(ut.Password) < 5 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, ut.Password, utf8.RuneCountInString(ut.Password), 5, true))
	}
	if utf8.RuneCountInString(ut.Password) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, ut.Password, utf8.RuneCountInString(ut.Password), 100, false))
	}
	return
}

// loginPayload user type.
type loginPayload struct {
	Email    *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the loginPayload type instance.
func (ut *loginPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 6, true))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 400, false))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 5, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 100, false))
		}
	}
	return
}

// Publicize creates LoginPayload from loginPayload
func (ut *loginPayload) Publicize() *LoginPayload {
	var pub LoginPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// LoginPayload user type.
type LoginPayload struct {
	Email    string `form:"email" json:"email" xml:"email"`
	Password string `form:"password" json:"password" xml:"password"`
}

// Validate validates the LoginPayload type instance.
func (ut *LoginPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, ut.Email, goa.FormatEmail, err2))
	}
	if utf8.RuneCountInString(ut.Email) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, ut.Email, utf8.RuneCountInString(ut.Email), 6, true))
	}
	if utf8.RuneCountInString(ut.Email) > 400 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, ut.Email, utf8.RuneCountInString(ut.Email), 400, false))
	}
	if utf8.RuneCountInString(ut.Password) < 5 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, ut.Password, utf8.RuneCountInString(ut.Password), 5, true))
	}
	if utf8.RuneCountInString(ut.Password) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.password`, ut.Password, utf8.RuneCountInString(ut.Password), 100, false))
	}
	return
}

// matchPayload user type.
type matchPayload struct {
	AwayGoals *int    `form:"away_goals,omitempty" json:"away_goals,omitempty" xml:"away_goals,omitempty"`
	AwayTeam  *string `form:"away_team,omitempty" json:"away_team,omitempty" xml:"away_team,omitempty"`
	AwayUser  *int    `form:"away_user,omitempty" json:"away_user,omitempty" xml:"away_user,omitempty"`
	HomeGoals *int    `form:"home_goals,omitempty" json:"home_goals,omitempty" xml:"home_goals,omitempty"`
	HomeTeam  *string `form:"home_team,omitempty" json:"home_team,omitempty" xml:"home_team,omitempty"`
	HomeUser  *int    `form:"home_user,omitempty" json:"home_user,omitempty" xml:"home_user,omitempty"`
}

// Validate validates the matchPayload type instance.
func (ut *matchPayload) Validate() (err error) {
	if ut.HomeTeam == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "home_team"))
	}
	if ut.AwayTeam == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "away_team"))
	}
	if ut.HomeGoals == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "home_goals"))
	}
	if ut.AwayGoals == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "away_goals"))
	}
	if ut.HomeUser == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "home_user"))
	}
	if ut.AwayUser == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "away_user"))
	}
	if ut.AwayGoals != nil {
		if *ut.AwayGoals < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.away_goals`, *ut.AwayGoals, 0, true))
		}
	}
	if ut.AwayTeam != nil {
		if utf8.RuneCountInString(*ut.AwayTeam) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.away_team`, *ut.AwayTeam, utf8.RuneCountInString(*ut.AwayTeam), 200, false))
		}
	}
	if ut.AwayUser != nil {
		if *ut.AwayUser < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.away_user`, *ut.AwayUser, 1, true))
		}
	}
	if ut.HomeGoals != nil {
		if *ut.HomeGoals < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.home_goals`, *ut.HomeGoals, 0, true))
		}
	}
	if ut.HomeTeam != nil {
		if utf8.RuneCountInString(*ut.HomeTeam) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.home_team`, *ut.HomeTeam, utf8.RuneCountInString(*ut.HomeTeam), 200, false))
		}
	}
	if ut.HomeUser != nil {
		if *ut.HomeUser < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.home_user`, *ut.HomeUser, 1, true))
		}
	}
	return
}

// Publicize creates MatchPayload from matchPayload
func (ut *matchPayload) Publicize() *MatchPayload {
	var pub MatchPayload
	if ut.AwayGoals != nil {
		pub.AwayGoals = *ut.AwayGoals
	}
	if ut.AwayTeam != nil {
		pub.AwayTeam = *ut.AwayTeam
	}
	if ut.AwayUser != nil {
		pub.AwayUser = *ut.AwayUser
	}
	if ut.HomeGoals != nil {
		pub.HomeGoals = *ut.HomeGoals
	}
	if ut.HomeTeam != nil {
		pub.HomeTeam = *ut.HomeTeam
	}
	if ut.HomeUser != nil {
		pub.HomeUser = *ut.HomeUser
	}
	return &pub
}

// MatchPayload user type.
type MatchPayload struct {
	AwayGoals int    `form:"away_goals" json:"away_goals" xml:"away_goals"`
	AwayTeam  string `form:"away_team" json:"away_team" xml:"away_team"`
	AwayUser  int    `form:"away_user" json:"away_user" xml:"away_user"`
	HomeGoals int    `form:"home_goals" json:"home_goals" xml:"home_goals"`
	HomeTeam  string `form:"home_team" json:"home_team" xml:"home_team"`
	HomeUser  int    `form:"home_user" json:"home_user" xml:"home_user"`
}

// Validate validates the MatchPayload type instance.
func (ut *MatchPayload) Validate() (err error) {
	if ut.HomeTeam == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "home_team"))
	}
	if ut.AwayTeam == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "away_team"))
	}

	if ut.AwayGoals < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.away_goals`, ut.AwayGoals, 0, true))
	}
	if utf8.RuneCountInString(ut.AwayTeam) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.away_team`, ut.AwayTeam, utf8.RuneCountInString(ut.AwayTeam), 200, false))
	}
	if ut.AwayUser < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.away_user`, ut.AwayUser, 1, true))
	}
	if ut.HomeGoals < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.home_goals`, ut.HomeGoals, 0, true))
	}
	if utf8.RuneCountInString(ut.HomeTeam) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.home_team`, ut.HomeTeam, utf8.RuneCountInString(ut.HomeTeam), 200, false))
	}
	if ut.HomeUser < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.home_user`, ut.HomeUser, 1, true))
	}
	return
}

// updateAccountPayload user type.
type updateAccountPayload struct {
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName  *string `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
}

// Validate validates the updateAccountPayload type instance.
func (ut *updateAccountPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 6, true))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 400, false))
		}
	}
	if ut.FirstName != nil {
		if utf8.RuneCountInString(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, utf8.RuneCountInString(*ut.FirstName), 2, true))
		}
	}
	if ut.FirstName != nil {
		if utf8.RuneCountInString(*ut.FirstName) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, utf8.RuneCountInString(*ut.FirstName), 200, false))
		}
	}
	if ut.LastName != nil {
		if utf8.RuneCountInString(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, utf8.RuneCountInString(*ut.LastName), 2, true))
		}
	}
	if ut.LastName != nil {
		if utf8.RuneCountInString(*ut.LastName) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, utf8.RuneCountInString(*ut.LastName), 200, false))
		}
	}
	return
}

// Publicize creates UpdateAccountPayload from updateAccountPayload
func (ut *updateAccountPayload) Publicize() *UpdateAccountPayload {
	var pub UpdateAccountPayload
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.FirstName != nil {
		pub.FirstName = ut.FirstName
	}
	if ut.LastName != nil {
		pub.LastName = ut.LastName
	}
	return &pub
}

// UpdateAccountPayload user type.
type UpdateAccountPayload struct {
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName  *string `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
}

// Validate validates the UpdateAccountPayload type instance.
func (ut *UpdateAccountPayload) Validate() (err error) {
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) < 6 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 6, true))
		}
	}
	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 400, false))
		}
	}
	if ut.FirstName != nil {
		if utf8.RuneCountInString(*ut.FirstName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, utf8.RuneCountInString(*ut.FirstName), 2, true))
		}
	}
	if ut.FirstName != nil {
		if utf8.RuneCountInString(*ut.FirstName) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.first_name`, *ut.FirstName, utf8.RuneCountInString(*ut.FirstName), 200, false))
		}
	}
	if ut.LastName != nil {
		if utf8.RuneCountInString(*ut.LastName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, utf8.RuneCountInString(*ut.LastName), 2, true))
		}
	}
	if ut.LastName != nil {
		if utf8.RuneCountInString(*ut.LastName) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.last_name`, *ut.LastName, utf8.RuneCountInString(*ut.LastName), 200, false))
		}
	}
	return
}

// updateMatchPayload user type.
type updateMatchPayload struct {
	AwayGoals *int    `form:"away_goals,omitempty" json:"away_goals,omitempty" xml:"away_goals,omitempty"`
	AwayTeam  *string `form:"away_team,omitempty" json:"away_team,omitempty" xml:"away_team,omitempty"`
	HomeGoals *int    `form:"home_goals,omitempty" json:"home_goals,omitempty" xml:"home_goals,omitempty"`
	HomeTeam  *string `form:"home_team,omitempty" json:"home_team,omitempty" xml:"home_team,omitempty"`
}

// Validate validates the updateMatchPayload type instance.
func (ut *updateMatchPayload) Validate() (err error) {
	if ut.AwayGoals != nil {
		if *ut.AwayGoals < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.away_goals`, *ut.AwayGoals, 0, true))
		}
	}
	if ut.AwayTeam != nil {
		if utf8.RuneCountInString(*ut.AwayTeam) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.away_team`, *ut.AwayTeam, utf8.RuneCountInString(*ut.AwayTeam), 200, false))
		}
	}
	if ut.HomeGoals != nil {
		if *ut.HomeGoals < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.home_goals`, *ut.HomeGoals, 0, true))
		}
	}
	if ut.HomeTeam != nil {
		if utf8.RuneCountInString(*ut.HomeTeam) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.home_team`, *ut.HomeTeam, utf8.RuneCountInString(*ut.HomeTeam), 200, false))
		}
	}
	return
}

// Publicize creates UpdateMatchPayload from updateMatchPayload
func (ut *updateMatchPayload) Publicize() *UpdateMatchPayload {
	var pub UpdateMatchPayload
	if ut.AwayGoals != nil {
		pub.AwayGoals = ut.AwayGoals
	}
	if ut.AwayTeam != nil {
		pub.AwayTeam = ut.AwayTeam
	}
	if ut.HomeGoals != nil {
		pub.HomeGoals = ut.HomeGoals
	}
	if ut.HomeTeam != nil {
		pub.HomeTeam = ut.HomeTeam
	}
	return &pub
}

// UpdateMatchPayload user type.
type UpdateMatchPayload struct {
	AwayGoals *int    `form:"away_goals,omitempty" json:"away_goals,omitempty" xml:"away_goals,omitempty"`
	AwayTeam  *string `form:"away_team,omitempty" json:"away_team,omitempty" xml:"away_team,omitempty"`
	HomeGoals *int    `form:"home_goals,omitempty" json:"home_goals,omitempty" xml:"home_goals,omitempty"`
	HomeTeam  *string `form:"home_team,omitempty" json:"home_team,omitempty" xml:"home_team,omitempty"`
}

// Validate validates the UpdateMatchPayload type instance.
func (ut *UpdateMatchPayload) Validate() (err error) {
	if ut.AwayGoals != nil {
		if *ut.AwayGoals < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.away_goals`, *ut.AwayGoals, 0, true))
		}
	}
	if ut.AwayTeam != nil {
		if utf8.RuneCountInString(*ut.AwayTeam) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.away_team`, *ut.AwayTeam, utf8.RuneCountInString(*ut.AwayTeam), 200, false))
		}
	}
	if ut.HomeGoals != nil {
		if *ut.HomeGoals < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.home_goals`, *ut.HomeGoals, 0, true))
		}
	}
	if ut.HomeTeam != nil {
		if utf8.RuneCountInString(*ut.HomeTeam) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.home_team`, *ut.HomeTeam, utf8.RuneCountInString(*ut.HomeTeam), 200, false))
		}
	}
	return
}