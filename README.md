# TILNotes API

TILNotes stands for Today I Learnt Notes. This is a note taking app of what you have learned today. So you can remember what you have learned.

## Stack

- Go
- GoFiber (Routing)
- GORM (ORM)
- PostgreSQL (DB)
- JWT (Authentication)

## How to Get Started

__*Make sure you have Go and PostgreSQL installed on your local machine.__

- Clone the project
- Use `go mod tidy` via the terminal to make sure you have the dependencies (you must be in the project root folder)
- Create `.env` file in the root folder with this following:
  ```
  DBHost=<insert_postgres_hostname> – defaults to: 127.0.0.1
  DBUser=<insert_postgres_username> - defaults to: postgres
  DBPassword=<insert_postgres_password> - defaults to: password (first time creation)
  DBName=<insert_database_name>
  DBPort=<insert_postgres_port> - defaults to: 5432

  JWTSecret=tilnotesapi
  ```
- Once you have the dependencies installed, to run the API server: type `go run main.go`

## Endpoints

### Authentication

- POST `/api/v1/users/auth/register`: Register a user

  ```
   curl --location --request POST 'http://localhost:1337/api/v1/users/auth/register' \
   --header 'Content-Type: application/json' \
   --data-raw '{
      "email": "test@test.com",
      "first_name": "test",
      "last_name": "test",
      "password": "test"
   }'
  ```

- POST `/api/v1/users/auth/login`: Login a user

  ```
  curl --location --request POST 'http://localhost:1337/api/v1/users/auth/login' \
  --header 'Content-Type: application/json' \
  --header 'Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg0MDgxMDYsImlzcyI6IjEifQ.y68iacn2rUzdsWGz3pTwq3U-ycw2-dPhWWSfRKOYdsU' \
  --data-raw '{
     "email": "test@test.com",
     "password": "test"
  }'
  ```

- GET `/api/v1/users/auth/current_user`: Retrieve current logged in user
  
  ```
  curl --location --request GET 'http://localhost:1337/api/v1/users/auth/current_user' \
  --header 'Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjgzNTU0NzgsImlzcyI6IjEifQ.m5N_TqjZ7vpSU-coN2m5YtOku-JeIEfv0nhlWsIidUA'
  ```

- POST `/api/v1/users/auth/logout`: Logout current user

  ```
  curl --location --request POST 'http://localhost:1337/api/v1/users/auth/logout'
  ```
  
---

### Notes

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
    "content": "Test 1",
    "category_id": 1,
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

### Categories

- GET `/api/v1/categories`: Retrieve all categories

  ```
  curl --location --request GET 'http://localhost:1337/api/v1/categories' \
  --header 'Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg0MDgxMDYsImlzcyI6IjEifQ.y68iacn2rUzdsWGz3pTwq3U-ycw2-dPhWWSfRKOYdsU'
  ```

- GET `/api/v1/categories/:id`: Retrieve a note with id param

  ```
  curl --location --request GET 'http://localhost:1337/api/v1/categories/1' \
  --header 'Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg0MDgxMDYsImlzcyI6IjEifQ.y68iacn2rUzdsWGz3pTwq3U-ycw2-dPhWWSfRKOYdsU'
  ```

- POST `/api/v1/categories`: Add a category

  ```
  curl --location --request POST 'http://localhost:1337/api/v1/categories' \
  --header 'Content-Type: application/json' \
  --header 'Cookie: jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg0MDgxMDYsImlzcyI6IjEifQ.y68iacn2rUzdsWGz3pTwq3U-ycw2-dPhWWSfRKOYdsU' \
  --data-raw '{
    "name": "GORM"
  }'
  ```

- PUT `/api/v1/categories/edit/:id`: Edit a category with id param

  ```
  curl -X PUT \
  'http://localhost:1337/api/v1/category/edit/1' \
  -H 'Accept: */*' \
  -H 'User-Agent: Thunder Client (https://www.thunderclient.io)' \
  -H 'Content-Type: application/json' \
  -d '{
    "title": "Test One",
    "content": "Test One"
  }'
  ```

- DELETE `/api/v1/categories/delete/:id`: Remove a category with id param

  ```
  curl -X DELETE \
  'http://localhost:1337/api/v1/notes/delete/1' \
  -H 'Accept: */*' \
  -H 'User-Agent: Thunder Client (https://www.thunderclient.io)'
  ```