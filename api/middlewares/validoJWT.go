package middlewares

import (
	http "net/http"

	pt "github.com/igorariza/Dockerized-Golang_API-MySql-React.js/internal/procesotoken"
)

// ValidoJWT permite validar JWT que retorna en la petición
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := pt.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}

}
