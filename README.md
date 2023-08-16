# go-rickandmorty-login API in Golang and Deploy to Heroku

This project demonstrates building a login API using the Go programming language and deploying it to the Heroku platform. The API allows users to log in and access "Rick and Morty" related data.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

## Features

- Build a login API in Go.
- Implement endpoints for user registration, login, and accessing "Rick and Morty" data.
- Utilize Heroku for deployment and hosting.

## Prerequisites

Before you begin, ensure you have the following installed:

- Go (https://golang.org/)
- Heroku CLI (https://devcenter.heroku.com/articles/heroku-cli)
- Database (https://app.planetscale.com/)

## Getting Started

1. Clone this repository:

   ```sh
   git clone https://github.com/igorariza/go-rickandmorty-login.git
   cd go-rickandmorty-login 
  ```
2. Install dependencies:

   ```sh
   go mod download
   ```
3. Create a `.env` file in the root directory of the project. Add the following environment variables

   ```sh
   PORT=8080
   DB_HOST=host
   DB_PORT=port
   DB_USER=user
   DB_PASS=password
   DB_NAME=database
   DB_SSL=disable
   ```
4. Run the application locally:

   ```sh
   go run main.go
   ```
5. Navigate to `http://localhost:8080/` or url heroku
URL for a dev server. Navigate to `https://go-rickandmorty-login-80476d397739.herokuapp.com/`

#### Get List characters

```http
  GET {{url}}/api/v1/characters
[
    {
        "ID": 1,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "name": "Rick Sanchez",
        "status": "Alive",
        "species": "Human",
        "gender": "Male",
        "image": "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
        "created": "2017-11-04T18:48:46.250Z"
    },
]
```
#### Create User

```http
  POST {{url}}/api/v1/users
{
    "name": "John Doe",
    "email": "johndoe@email.com",
    "password": "admin",
    "address": "Street 123",
    "birthdate": "1990-01-01",
    "city": "The world"
}

RESPONSE
{
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "name": "John Doe",
    "email": "johndoe@email.com",
    "password": "",
    "address": "Street 123",
    "birthdate": "1990-01-01",
    "city": "The world"
}
```

#### Login

```http
  POST {{url}}/api/v1/login
{
    "email": "johndoe@email.com",
	  "password": "admin"
}

RESPONSE
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZW1haWwuY29tIiwiZXhwIjoxNjkyMjY2MjMxLCJwYXNzd29yZCI6IiJ9.8fiIJW5htHKNQiXRk-Ul1xAingYawXZDBPYWMFXnfuQ"
}
{
    "success": true,
    "status": 200,
    "result": {
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "email": "johndoe@email.com",
    }
}
```