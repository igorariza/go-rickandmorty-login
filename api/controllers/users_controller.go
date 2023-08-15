package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/igorariza/go-rickandmorty-login/api/models"
	"github.com/igorariza/go-rickandmorty-login/api/responses"
	jwt "github.com/igorariza/go-rickandmorty-login/internal/jwt"
	"github.com/igorariza/go-rickandmorty-login/internal/web"
)

func (server *Server) LoginUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	user := models.User{}
	body := r.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.LoginUser(server.DB)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(userCreated.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(*userCreated)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar general el Token correspondiente "+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

	//Grabacion de cookie en el lado Usuario
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
	web.Success(&models.User{
		Email: userCreated.Email,
	}, http.StatusOK).Send(w)

}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI))
	responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) GetUserId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAllusers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}
