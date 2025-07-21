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

## Setup
1. Clone this repo
2. Run `go mod tidy`
3. Run `go test ./...` to verify everything is OK
4. Run `go run cmd/api/main.go` to start the server on port :8080
5. Use the provided Postman collection in `docs/` to test endpoints

## API Endpoints
- `POST /publish` – Publish a new message
- `POST /follow` – Follow a new user
- `GET /usertimeline` – Get the timeline for a user

## Example Requests
### Publish a Message
```
curl -X POST http://localhost:8080/publish -d '{"user_id": "alice", "content": "Hello world!"}' -H 'Content-Type: application/json'
```

### Follow a User
```
curl -X POST http://localhost:8080/follow -d '{"user_id": "alice", "new_follow": "bob"}' -H 'Content-Type: application/json'
```

### Get User Timeline
```
curl -X GET http://localhost:8080/usertimeline -d '{"user_id": "alice"}' -H 'Content-Type: application/json'
```

## Notes
- When running locally with the in-memory repository, restart the server to reset the database.
- See `docs/microblogging-platform.postman_collection.json` for more request examples.

---

For questions or contributions, open an issue or pull request.


