# TILNotes API

TILNotes stands for Today I Learnt Notes. This is a note taking app of what you have learned today. So you can remember what you have learned.

## Stack

- Go
- GoFiber (Routing)
- GORM (ORM)
- PostgreSQL (DB)

## How to Get Started

- Clone the project
- Use `go mod tidy` via the terminal to make sure you have the dependencies (you must be in the project root folder)
- Once you have the dependencies installed, to run the API server: type `go run main.go`

## Endpoints

- GET `/api/v1/notes`: Retrieve all notes
  
  ```
  curl -X GET \
  'http://localhost:1337/api/v1/notes' \
  -H 'Accept: */*' \
  -H 'User-Agent: Thunder Client (https://www.thunderclient.io)'
  ```
  
- GET `/api/v1/notes/:id`: Retrieve a note with id param

  ```
  curl -X GET \
  'http://localhost:1337/api/v1/notes/1' \
  -H 'Accept: */*' \
  -H 'User-Agent: Thunder Client (https://www.thunderclient.io)'
  
  ```

- POST `/api/v1/notes`: Add a note

  ```
  curl -X POST \
  'http://localhost:1337/api/v1/notes' \
  -H 'Accept: */*' \
  -H 'User-Agent: Thunder Client (https://www.thunderclient.io)' \
  -H 'Content-Type: application/json' \
  -d '{
    "title": "Test 1",
    "content": "Test 1"
  }'
  ```

- PUT `/api/v1/notes/edit/:id`: Edit a note with id param

  ```
  curl -X PUT \
  'http://localhost:1337/api/v1/notes/edit/1' \
  -H 'Accept: */*' \
  -H 'User-Agent: Thunder Client (https://www.thunderclient.io)' \
  -H 'Content-Type: application/json' \
  -d '{
    "title": "Test One",
    "content": "Test One"
  }'
  ```

- DELETE `/api/v1/notes/delete/:id`: Remove a note with id param

  ```
  curl -X DELETE \
  'http://localhost:1337/api/v1/notes/delete/1' \
  -H 'Accept: */*' \
  -H 'User-Agent: Thunder Client (https://www.thunderclient.io)'
  ```
