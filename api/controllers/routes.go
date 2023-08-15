package controllers

import "github.com/igorariza/Dockerized-Golang_API-MySql-React.js/api/middlewares"

func (s *Server) initializeRoutes() {

	// Character Route
	s.Router.HandleFunc("/api/v1/characters", middlewares.SetMiddlewareJSON(s.GetAllCharacters)).Methods("GET")
	s.Router.HandleFunc("/api/v1/character", middlewares.SetMiddlewareJSON(s.CreateCharacter)).Methods("POST")
	s.Router.HandleFunc("/api/v1/character/{id}", middlewares.SetMiddlewareJSON(s.GetCharacterId)).Methods("GET")

	//Users routes
	s.Router.HandleFunc("/api/v1/user", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/v1/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/v1/login", middlewares.SetMiddlewareJSON(s.LoginUsersHandler)).Methods("POST")

}
