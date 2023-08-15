package internal

import (
	"errors"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	usr "github.com/igorariza/go-rickandmorty-login/api/models"
)

var Email string
var IDUsuario string

// ProcesoToken procesa token para obtener sus valores
func ProcesoToken(tk string) (*usr.Claim, bool, string, error) {
	claims := &usr.Claim{}
	miClave := []byte(os.Getenv("MY_CLAVE"))

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		Email = claims.Email
		IDUsuario = claims.ID
		return claims, true, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
