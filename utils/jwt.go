package utils

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/rymccue/fifa-tracker/helpers"
	uuid "github.com/satori/go.uuid"
)

func GenerateJWTToken(id int) (string, error) {
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	exp := time.Now().Add(time.Duration(24*30) * time.Hour).Unix()
	token.Claims = jwtgo.MapClaims{
		"iss": "Issuer",              // who creates the token and signs it
		"aud": "Audience",            // to whom the token is intended to be sent
		"exp": exp,                   // time when the token will expire (10 minutes from now)
		"jti": uuid.NewV4().String(), // a unique identifier for the token
		"iat": time.Now().Unix(),     // when the token was issued/created (now)
		"sub": "subject",             // the subject/principal is whom the token is about
		"id":  id,
	}
	pk, err := helpers.LoadJWTPrivateKey()
	if err != nil {
		return "", err
	}
	return token.SignedString(pk)
}
