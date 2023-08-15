package internal

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	users "github.com/igorariza/go-rickandmorty-login/api/models"
)

/*GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t users.LoginUser) (string, error) {

	miClave := []byte(os.Getenv("API_SECRET"))

	payload := jwt.MapClaims{
		"email":    t.Email,
		"password": t.Password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
