// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "FIFA Tracker": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/rymccue/fifa-tracker/design
// --out=$(GOPATH)/src/github.com/rymccue/fifa-tracker
// --version=v1.2.0-dirty

package app

import (
	"fmt"
	"strings"
)

// AccountHref returns the resource href.
func AccountHref(accountID interface{}) string {
	paramaccountID := strings.TrimLeftFunc(fmt.Sprintf("%v", accountID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/account/%v", paramaccountID)
}

// MatchHref returns the resource href.
func MatchHref(matchID interface{}) string {
	parammatchID := strings.TrimLeftFunc(fmt.Sprintf("%v", matchID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/match/%v", parammatchID)
}
