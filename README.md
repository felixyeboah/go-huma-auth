# Go REST API with Huma and Chi

## Overview

This project is a RESTful API built using Go, leveraging the Huma and Chi libraries for routing and handling HTTP requests. The API implements a basic authentication system with role-based access control, where users can have different privileges based on their roles (e.g., `user`, `admin`). The project also integrates SQLC for generating type-safe database queries.

## Table of Contents

- [Installation](#installation)
- [Project Structure](#project-structure)
- [Database Schema](#database-schema)
- [API Endpoints](#api-endpoints)
- [Authentication & Authorization](#authentication--authorization)
- [Testing](#testing)
- [Learning Journey](#learning-journey)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/felixyeboah/go-huma-auth.git
   cd go-huma-auth
   ```

2. **Set Up the Database:**

Make sure you have PostgreSQL installed.
Create a database and run the migrations located in the migrations/ directory.

3. **Install Dependencies:*

```bash
  go mod tidy
```

4. **Run the Application:**

```bash
go run cmd/server/main.go
```

## Project Structure

```bash
go-huma-auth/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── auth/
│   │   ├── handler.go
│   │   ├── repository.go
│   │   ├── service.go
│   └── middleware/
│       └── auth.go
├── migrations/
│   └── 001_create_users_table.sql
├── pkg/
│   ├── db/
│   │   ├── db.go
│   └── utils/
│       └── utils.go
├── sql
│   ├── queries/
│   ├── sqlc/
├── go.mod
├── go.sum
├── Makefile
└── sqlc.yaml
```

## Database Schema

The database schema includes tables for users, roles, privileges, and the associations between them. The roles table defines roles like user and admin, while the privileges table defines various permissions that can be assigned to these roles.

## API Endpoints
###  Authentication

- POST /register: Register a new user.
- POST /login: Log in with email and password.

### Authorization

- GET /profile: Get the user's profile (requires authentication).

## Authentication & Authorization

This API uses role-based access control (RBAC) to manage access to various endpoints. Users are assigned roles, and each role has specific privileges. The privileges determine what actions a user can perform in the system.

## Testing

Testing is implemented using stretchr/testify. Unit tests cover the repository, service, and handler layers.

To run tests:

```bash
go test ./...
```

## Learning Journey

This project is part of my journey to learn Go and build robust, scalable APIs. Through this project, I am gaining a deep understanding of Go's ecosystem, including:

- Goroutines and Concurrency: Leveraging Go's concurrency model for handling multiple requests.
- Error Handling: Understanding Go's approach to error handling.
- Middleware: Implementing custom middleware for authentication and authorization.
- Testing: Writing effective unit and integration tests.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have suggestions or improvements.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

```bash

You can copy this directly into your `README.md` file.
```