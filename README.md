# Events Service

This is a Go project that provides a service for managing user subscriptions to events.

<div >
	<code><img width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original-wordmark.svg" alt="Go" title="Go"/></code>
	<code><img width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original.svg" alt="PostgreSQL" title="PostgreSQL"/></code>
	<code><img width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/sentry/sentry-original.svg" alt="Sentry" title="Sentry"/></code>
	<code><img width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/grpc/grpc-original.svg" alt="gRPC" title="gRPC"/></code>
	<code><img width="50" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/datadog/datadog-original.svg" alt="Datadog" title="Datadog"/></code>
	<code><img width="50" src="https://raw.githubusercontent.com/marwin1991/profile-technology-icons/refs/heads/main/icons/gcp.png" alt="GCP" title="GCP"/></code>
	<code><img width="50" src="https://raw.githubusercontent.com/marwin1991/profile-technology-icons/refs/heads/main/icons/rabbitmq.png" alt="RabbitMQ" title="RabbitMQ"/></code>
</div>

## Getting Started

To get started with this project, you'll need to have Go installed on your machine. You'll also need to have a running instance of PostgreSQL and the users-service running.

### Prerequisites

- Go
- PostgreSQL
- Docker (optional)
- Users Service (gRPC)

### Installation

1. Clone the repository:
   ```sh
   git clone gitlab.com/velo-company/services/events-service
   ```
2. Install the dependencies:
   ```sh
   go mod tidy
   ```
3. Set up the environment variables:
   ```sh
   cp .env.example .env
   ```
   The following environment variables are used:
   - `POSTGRES_CONNECTION_STRING`: The connection string for the PostgreSQL database.
   - `RSA_PUBLIC_KEY`: The path to the RSA public key for verifying JWT tokens.
   - `GEMINI_API_KEY`: The API key for the Gemini service.
   - `USER_SERVICE_GRPC_ADDRESS`: The address for the User Service gRPC server.

4. Run the database migrations:
   ```sh
   # Assuming makefile exists or using a migration tool
   migrate -path db/migrations -database "postgres://user:password@host:port/dbname?sslmode=disable" up
   ```
5. Run the server:
   ```sh
   go run cmd/api/main.go
   ```

### Docker

To build and run the project using Docker, you can use the following commands:

1. Build the Docker image:
   ```sh
   docker build -t events-service .
   ```
2. Run the Docker container:
   ```sh
   docker run --env-file .env -p 8080:8080 events-service
   ```

## Folder Structure

The project is organized into the following folders:

- `cmd/api`: Contains the main application entry point.
- `db/migrations`: Contains the database migrations.
- `internal`: Contains the core application logic.
  - `adapters`: Contains the adapters for connecting to external services.
    - `ai`: Contains adapters for AI services like Gemini.
    - `database`: Contains the adapters for modifying PostgreSQL Data.
    - `grpc`: Contains the adapters for calling procedures in other APIs.
    - `http`: Contains the REST controllers (handlers).
  - `core`: Contains the core domain logic.
    - `entities`: Contains the domain models.
    - `errors`: Contains custom error types.
    - `ports`: Contains the interfaces for the repositories and services.
    - `services`: Contains the application services.
- `proto`: Contains the protobuf files for the gRPC services.

## API Endpoints

All endpoints are prefixed with `/api/events/v1` and require JWT authentication.

- `POST /subscribe/:id`: Subscribes the authenticated user to the event with the specified ID.
- `POST /cancel-subscription/:id`: Cancels the authenticated user's subscription to the event with the specified ID.
- `POST /confirm-subscription/:id`: Confirms the user's subscription to the event with a confirmation code.
- `GET /confirmation-code/:id`: Retrieves the confirmation code for a user's subscription to an event.

## Development

### Creating Migrations

To create a new migration, you can use a tool like `migrate`:

```sh
migrate create -ext sql -dir db/migrations -seq create_my_migration
```

This will create new up and down migration files in the `db/migrations` folder.

### Running Migrations

To run the migrations, you can use the following command:

```sh
migrate -path db/migrations -database "postgres://user:password@host:port/dbname?sslmode=disable" up
```

Make sure to replace the connection string with your own.


### Generating Protobuf

To generate the `.pb.go` file from the `.proto` file, you can use the following command:

```sh
protoc --go_out=. --go-grpc_out=. proto/user.proto
```

### Swagger Documentation

To generate the swagger documentation, you can use the following command:

```sh
swag init -g cmd/api/main.go -o docs/
```

This will generate the `docs/` folder with the swagger documentation.

## Dependencies

The project uses the following primary dependencies:

- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)
- [github.com/joho/godotenv](https://github.com/joho/godotenv)
- [github.com/lib/pq](https://github.com/lib/pq)
- [google.golang.org/grpc](https://grpc.io/)
- [google.golang.org/protobuf](https://developers.google.com/protocol-buffers)
- [github.com/google/generative-ai-go](https://github.com/google/generative-ai-go)

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
