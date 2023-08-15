# go-rickandmorty-login

## Deployment

To deploy this project run

```bash
  docker-compose up 
```
for a dev server. Navigate to `http://localhost:8080/`
## API Reference

#### Get List characters

```http
  GET /api/v1/characters
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
#### Get item

```http
  GET /get-beer/${id}
```