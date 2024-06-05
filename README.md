
```
   ___  ____  ____  _______  __
  / _ \/ __ \/ __ \/_  __/ |/_/
 / , _/ /_/ / /_/ / / / _>  <  
/_/|_|\____/\____/ /_/ /_/|_|  
                               
```

                               



🚀 Go Starter with Postgres + Go + Swagger Default Auth + REST API + Cache Middleware + Goroutines 🚀

A comprehensive starter kit for building RESTful APIs in Go with PostgreSQL as the database backend. This starter kit includes Swagger for API documentation, default authentication setup, dynamic tooling for database migrations and CRUD API generation, cache middleware for performance optimization, and goroutines for concurrent execution. Kickstart your Go project with ease and efficiency!

Features:
- PostgreSQL database integration
- RESTful API endpoints
- Swagger documentation
- Default authentication setup
- Dynamic tooling for database migrations and CRUD API generation
- Cache middleware for performance optimization
- Goroutines for concurrent execution



## Getting Started

### Prerequisites

- Go 1.16 or higher
- Docker
- Docker Compose
- Node.js

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/JubaerHossain/rootx.git
    cd rootx
    ```

2. Build and run the application using Docker Compose:
    ```sh
    docker-compose up --build
    ```

3. Access the application at `http://localhost:8080`.

### Usage

- To start the server:
    ```sh
    go run cmd/main.go
    ```

- To run tests:
    ```sh
    go test ./...
    ```


## install dependencies

```bash
make install
```

## run the project [development mode]

```bash
make dev
```

```bash
Select a command:
1. Create Migration
2. Create Migration with Seeder
3. Apply Migrations
4. Run Seeders
Enter the command number: 4
Running seeders...
2024/06/04 19:35:32 seeds files executed successfully
Task completed successfully
```


## create a new migration example [select 1]

```bash
make rootx 
```
## run the migration [select 2]

```bash
make rootx

```

## seed the database [select 3]

```bash
make rootx
```
## create a new module

```bash
make command name=module_name
```

## build the project

```bash
make build
```

## run the project [production mode]

```bash
make run
```

### uses database postgres

```bash
sudo -u postgres psql postgres
```

### Code check by golangci-lint [linting and formatting code](https://golangci-lint.run/welcome/quick-start/)

```bash
make lint
```

### docker volume list

```bash
docker volume ls
```

## docker volume remove

```bash
docker volume rm <volume_name>
```

## docker shell

```bash
docker exec -it <container_id> /bin/sh
```

or

```bash
docker exec -it <container_id> bash
```

## docker postgres

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres
```

## docker postgres with volume

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 -v postgres:/var/lib/postgresql/data postgres
```

## docker postgres with volume and password

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 -v postgres:/var/lib/postgresql/data postgres
```

## docker postgres with volume and password and user

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d -p 5432:5432 -v postgres:/var/lib/postgresql/data postgres
```

## docker postgres with volume and password and user and database

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=restaurant-api -d -p 5432:5432 -v postgres:/var/lib/postgresql/data postgres
```

## postgres shell

```bash
psql -h localhost -U postgres
```

## Directory Structure

```
project
│
.
├── cmd
│   └── main.go
├── config
│   └── config.go
├── internal
│   ├── app.go
│   ├── domain
│   │   ├── comment
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   ├── post
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   ├── user
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   ├── interfaces
│   │   ├── grpc
│   │   │   ├── proto
│   │   │   │   └── service.proto
│   │   │   └── server.go
│   │   ├── http
│   │   │   ├── handler.go
│   │   │   ├── router.go
│   │   │   └── server.go
│   ├── usecase
│   │   ├── comment.go
│   │   ├── post.go
│   │   └── user.go
│   └── templates
│       ├── layouts
│       │   └── base.html
│       ├── partials
│       │   ├── footer.html
│       │   ├── header.html
│       │   └── sidebar.html
│       └── views
│           ├── auth
│           │   ├── login.html
│           │   └── register.html
│           ├── cart
│           │   └── index.html
│           ├── products
│           │   ├── detail.html
│           │   └── list.html
│           └── home.html
├── pkg
│   ├── auth
│   │   └── jwt.go
│   ├── cache
│   │   └── redis.go
│   ├── db
│   │   └── postgres.go
│   ├── logger
│   │   └── logger.go
│   ├── middleware
│   │   └── auth.go
│   ├── response
│   │   └── response.go
│   ├── tasks
│   │   └── background.go
│   └── workerpool
│       └── workerpool.go
├── static
│   ├── css
│   │   └── tailwind.css
│   ├── js
│   │   └── main.js
│   └── images
└── go.mod

```

## Description

- **cmd**: Entry point of the application.
  - `main.go`: Main executable file.

- **config**: Configuration settings.
  - `config.go`: Configuration management.

- **pkg**: External packages.
  - **cache**: Caching mechanisms.
    - `redis.go`: Redis cache implementation.
  - **db**: Database connections.
    - `postgres.go`: PostgreSQL database implementation.
  - **logger**: Logging utilities.
    - `logger.go`: Logger setup.
  - **middleware**: HTTP middleware.
    - `auth.go`: Authentication middleware.
    - `cors.go`: CORS handling.
    - `logging.go`: Request logging.
    - `ratelimit.go`: Rate limiting.
  - **auth**: Authentication mechanisms.
    - `jwt.go`: JWT handling.
  - **response**: HTTP response utilities.
    - `response.go`: Response formatting.
  - **tasks**: Background tasks.
    - `background.go`: Background task management.
  - **workerpool**: Worker pool for concurrent task processing.
    - `workerpool.go`: Worker pool implementation.

- **internal**: Internal application logic.
  - **domain**: Domain logic for entities.
    - **user**: User domain.
      - `entity.go`: User entity definition.
      - `repository.go`: User repository interface.
      - `service.go`: User service implementation.
    - **post**: Post domain.
      - `entity.go`: Post entity definition.
      - `repository.go`: Post repository interface.
      - `service.go`: Post service implementation.
    - **comment**: Comment domain.
      - `entity.go`: Comment entity definition.
      - `repository.go`: Comment repository interface.
      - `service.go`: Comment service implementation.
  - **interfaces**: External interfaces (e.g., HTTP, gRPC).
    - **grpc**: gRPC server and proto definitions.
      - `server.go`: gRPC server implementation.
      - **proto**: Protocol buffer definitions.
        - `service.proto`: gRPC service definitions.
    - **http**: HTTP server.
      - `handler.go`: HTTP handlers.
      - `router.go`: HTTP routing.
      - `server.go`: HTTP server setup.
  - **usecase**: Application use cases.
    - `user.go`: User use cases.
    - `post.go`: Post use cases.
    - `comment.go`: Comment use cases.
  - `app.go`: Application setup and initialization.

- **html**: Frontend templates and static files.
  - `index.html`: Homepage template.
  - `login.html`: Login page template.
  - `posts.html`: Posts listing page template.
  - `post.html`: Single post page template.
  - `styles.css`: CSS styles.
  - `htmx.min.js`: HTMX library.

- **Dockerfile**: Docker configuration for containerizing the application.
- **docker-compose.yml**: Docker Compose configuration for multi-container applications.
- **go.mod**: Go module file.
- **go.sum**: Go dependencies.
- **package.json**: Node.js package configuration.
- **postcss.config.js**: PostCSS configuration.
- **tailwind.config.js**: Tailwind CSS configuration.


## Features

- [x] Golang
- [x] DDD
- [x] Clean Architecture
- [x] Docker
- [x] Makefile
- [x] Swagger
- [x] Gorm
- [x] JWT
- [x] Viper
- [x] Logger[Zap]
- [x] Unit Test

## Contributing

1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [HTMX](https://htmx.org)
- [Tailwind CSS](https://tailwindcss.com)
- [Go](https://golang.org)
- [Docker](https://www.docker.com)

---

Copy and paste the above markdown text into your `README.md` file on GitHub. This file provides an organized overview of your project, including the directory structure, descriptions of each component, installation instructions, usage guidelines, contribution steps, and acknowledgements.