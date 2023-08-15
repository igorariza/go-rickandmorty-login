package models

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
)

// Claim es la estrucutura usada para procesar el jwt
type Claim struct {
	gorm.Model
	Email string `json:"email"`
	ID	string   `json:"id"`
	jwt.ClaimsValidator
}
