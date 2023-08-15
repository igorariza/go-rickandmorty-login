package controllers

import (
	"net/http"

	"github.com/igorariza/Dockerized-Golang_API-MySql-React.js/api/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/igorariza/Dockerized-Golang_API-MySql-React.js/api/models"
	"github.com/igorariza/Dockerized-Golang_API-MySql-React.js/api/utils/formaterror"
)


func (server *Server) CreateCharacter(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	character := models.Character{}
	err = json.Unmarshal(body, &character)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	character.Prepare()
	err = character.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	characterCreated, err := character.SaveCharacter(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, characterCreated.ID))
	responses.JSON(w, http.StatusCreated, characterCreated)
}

func (server *Server) GetCharacterId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	character := models.Character{}
	characterGotten, err := character.FindCharacterByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, characterGotten)
}

func (server *Server) GetAllCharacters(w http.ResponseWriter, r *http.Request) {

	character := models.Character{}

	characters, err := character.FindAllcharacters(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, characters)
}
