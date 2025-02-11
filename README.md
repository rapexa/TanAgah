# TanAgah

This is a Training and nutrition programming system API built with the Gin framework in Go.

## Features

- User management (Create, Read, Update, Delete)
- Versioned API
- Clean architecture
- MySQL database with GORM
- Logger system

## Setup

1. Clone the repository.
2. Create a `.env` file based on the `.env.example`.
3. Run `go mod tidy` to install dependencies.
4. Run `go run cmd/main.go` to start the server.

## API Endpoints

- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users/:id` - Get a user by ID
- `PUT /api/v1/users/:id` - Update a user by ID
- `DELETE /api/v1/users/:id` - Delete a user by ID

## Database

Run the `database.sql` script to create the necessary database and table.

## License

MIT