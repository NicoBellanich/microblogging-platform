# Microblogging-platform

A simplified Twitter-like microblogging platform built in Go. Users can post messages, follow others, and view a personalized timeline. The project demonstrates clean architecture, modular design, and in-memory or pluggable persistence.

## Features
- Publish short messages (up to 280 characters)
- Follow other users
- View a timeline of messages from users you follow
- Simple REST API
- Easily switch between in-memory, test, and production repositories

## Architecture Overview
- **cmd/api/main.go**: Entry point for the API server
- **cmd/server/**: Server setup and dependency wiring
- **internal/controllers/**: HTTP handlers/controllers
- **internal/usecase/**: Application business logic (use cases)
- **internal/domain/**: Core domain models
- **internal/platform/repository/**: Repository interfaces and implementations (in-memory, prod, test)
- **config/**: Environment and configuration management

## Setup with terminal
1. Clone this repo
2. Run `go mod tidy`
3. Run `go test ./...` to verify everything is OK
4. Run `go run cmd/api/main.go` to start the server on port :8080
5. Use the provided Postman collection in `docs/` to test endpoints

## Docker Setup

You can easily run the application using Docker. Make sure you have Docker installed on your system.

### 1. Build the image

run this to download current dockerhub image

```bash
docker pull nicolasbellanich/microblogging-platform-final:v1
```

### 2. Run the container

- run

```bash
docker run -p 8080:8080 nicolasbellanich/microblogging-platform-final:v1
```

This will start the API at [http://localhost:8080](http://localhost:8080).

You can test the endpoints using the provided Postman collection in `docs/` or with the `curl` examples above.

## API Endpoints

### Create User
- **POST** `/user/create`
- **Body:**
```json
{
  "user_id": "alice"
}
```
- **Response:**
```json
{
  "message": "User created successfully",
  "status": 201
}
```

### Get User by Username
- **GET** `/user`
- **Body:**
```json
{
  "user_id": "alice"
}
```
- **Response:**
```json
{
  "username": "alice",
  "following": ["bob"],
  "publications": ["Hello world!", ...]
}
```

### Get User Timeline
- **GET** `/user/timeline`
- **Body:**
```json
{
  "user_id": "alice"
}
```
- **Response:**
```json
{
  "feed": [
    {
      "id": "...",
      "user_id": "bob",
      "content": "Hello from bob!",
      "created_at": "2024-06-01T12:00:00Z"
    },
    ...
  ]
}
```

### Publish a Message
- **POST** `/user/publish`
- **Body:**
```json
{
  "user_id": "alice",
  "content": "Hello world!"
}
```
- **Response:**
```json
{
  "message": "Message published successfully",
  "status": 201
}
```

### Follow a User
- **POST** `/user/following`
- **Body:**
```json
{
  "user_id": "alice",
  "new_follow": "bob"
}
```
- **Response:**
```json
{
  "message": "New Follower added successfully",
  "status": 201
}
```

## Example Requests

### Create a User
```
curl -X POST http://localhost:8080/user/create -d '{"user_id": "alice"}' -H 'Content-Type: application/json'
```

### Get User by Username
```
curl -X GET http://localhost:8080/user -d '{"user_id": "alice"}' -H 'Content-Type: application/json'
```

### Publish a Message
```
curl -X POST http://localhost:8080/user/publish -d '{"user_id": "alice", "content": "Hello world!"}' -H 'Content-Type: application/json'
```

### Follow a User
```
curl -X POST http://localhost:8080/user/following -d '{"user_id": "alice", "new_follow": "bob"}' -H 'Content-Type: application/json'
```

### Get User Timeline
```
curl -X GET http://localhost:8080/user/timeline -d '{"user_id": "alice"}' -H 'Content-Type: application/json'
```

## Notes
- When running locally with the in-memory repository, restart the server to reset the database.
- See `docs/microblogging-platform-v2.postman_collection.json` for more request examples.

## Visualizar la documentación Swagger/OpenAPI localmente

You can run Swagger UI easily using Docker to visualize API documentation:

```bash
docker run --rm -p 8081:8080 -v "$(pwd)/docs/openapi.yaml:/openapi.yaml" -e SWAGGER_JSON=/openapi.yaml swaggerapi/swagger-ui
```

Then open your browser [http://localhost:8081](http://localhost:8081) to see and test endpoints.
