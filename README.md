# TILNotes API

TILNotes stands for Today I Learnt Notes. This is a note taking app of what you have learned today. So you can remember what you have learned.

## Stack

- Go
- GoFiber (Routing)
- GORM (ORM)
- PostgreSQL (DB)
- JWT (Authentication)

## How to Get Started

### Run on localhost:

__*Make sure you have Go and PostgreSQL installed on your host machine.__

- Clone the project
  - via ssh: `git clone git@github.com:mofodox/TILNotes.git`
  - via https: `git clone https://github.com/mofodox/TILNotes.git`
- Use `go mod download` via the terminal to make sure you have the dependencies (you must be in the project root folder)
- Rename the `.env.sample` to `.env` in the root folder with this following:
  ```
  DBHost=<insert_postgres_hostname> – defaults to: 127.0.0.1
  DBUser=<insert_postgres_username> - defaults to: postgres
  DBPassword=<insert_postgres_password> - defaults to: password (first time creation)
  DBName=<insert_database_name>
  DBPort=<insert_postgres_port> - defaults to: 5432

  JWTSecret=tilnotesapi
  ```
- Once you have the dependencies installed, to run the API server: type `go run main.go`
- To test the server is working, visit `http://localhost:1337/` and you can see a string `Hello World` is sent.

### Run on Docker Container:

__*Make sure you have Go and Docker installed on your host machine.__

- Clone the project
  - via ssh: `git clone git@github.com:mofodox/TILNotes.git`
  - via https: `git clone https://github.com/mofodox/TILNotes.git`
- Build the custom image: `docker build -t <any_name_you_want> .`
- After the image has been successfully built: `docker-compose -up`
- To test the server is working, visit `http:localhost:8080` and you can see a string `Hello World` is sent.

You can play around with TILNotesAPI with a REST API client such as POSTMAN or whichever you prefer.

Have fun!

---

## Endpoints

### Authentication

JWT will be created when the user login, and the jwt token is saved in cookie. If you look in the `router` file,
we created a fiber middleware `middlewares.AuthRequired` to handle the authorization on the endpoints.

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
  
---

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

---

## Contributing

TILNotesAPI is an open source project. Please feel free to contribute and when contributing, please follow the Code of Conduct.

### Issues

Feel free to submit issues and enhancement requests.

### How to Get Started

In general we follow the "fork-and-pull" Git workflow.

1. Fork the repo on GitHub 
2. Clone the project to your own machine
3. Commit changes to your own branch
4. Push your work back up to your fork
5. Submit a Pull request so that we can review your changes

NOTE: Be sure to merge the latest from "main" branch before making a pull request.

---

## Contributors

- [Khairul Akmal](http://github.com/mofodox)