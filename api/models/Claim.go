package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claim is the token payload
type Claim struct {
	User User `json:"user"`
	jwt.StandardClaims
}
