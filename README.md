# rootx



project
│
├── cmd
│   └── main.go
├── config
│   └── config.go
├── pkg
│   ├── cache
│   │   └── redis.go
│   ├── db
│   │   └── postgres.go
│   ├── logger
│   │   └── logger.go
│   ├── middleware
│   │   ├── auth.go
│   │   ├── cors.go
│   │   ├── logging.go
│   │   └── ratelimit.go
│   ├── auth
│   │   └── jwt.go
│   ├── response
│   │   └── response.go
│   ├── tasks
│   │   └── background.go
│   └── workerpool
│       └── workerpool.go
├── internal
│   ├── domain
│   │   ├── user
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   ├── post
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── comment
│   │       ├── entity.go
│   │       ├── repository.go
│   │       └── service.go
│   ├── interfaces
│   │   ├── grpc
│   │   │   ├── server.go
│   │   │   └── proto
│   │   │       └── service.proto
│   │   └── http
│   │       ├── handler.go
│   │       ├── router.go
│   │       └── server.go
│   ├── usecase
│   │   ├── user.go
│   │   ├── post.go
│   │   └── comment.go
│   └── app.go
├── html
│   ├── index.html
│   ├── login.html
│   ├── posts.html
│   ├── post.html
│   ├── styles.css
│   └── htmx.min.js
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── package.json
├── postcss.config.js
└── tailwind.config.js

