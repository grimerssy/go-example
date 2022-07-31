# Example go project

---

### Introduction

This is a server application built using [go-template](https://github.com/grimerssy/go-template) 
to showcase how to use it.

You might find this project over-engineered as it is just a simple API which 
lets 
you signup, login and get a greeting from it. Yet it uses three-layer 
architecture and can be called using gRPC, RESTful HTTP/1.1 and HTTP/2.

Why is it like that? It is, as the name suggests, an example. Its goal is to 
not only show how to use [go-template](https://github.com/grimerssy/go-template),
but to show my approach to building real-world backend apps using Go, which 
you might find useful.

The project follows [standard Go project layout](https://github.com/golang-standards/project-layout)
and is directly written to be a guide to writing more maintainable and 
easily extensible code in Go. 

___

### Dependencies

Let's move on to libraries, which I used in this project, and which this 
application is build around. 

- [grpc-go](https://github.com/grpc/grpc-go) - building gRPC server
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) - 
  building RESTful HTTP proxy on top of gRPC server
- [wire](https://github.com/google/wire) - dependency injection
- [migrate](https://github.com/golang-migrate/migrate) - working with 
  database migrations
- [viper](https://github.com/spf13/viper) - reading configuration files
- [jwt-go](https://github.com/golang-jwt/jwt) - JWT authorization
- [zap](https://github.com/uber-go/zap) - logging
- [optimus-go](https://github.com/pjebs/optimus-go) - hiding database ids from 
end user
- [sqlx](https://github.com/jmoiron/sqlx) - better SQL experience
- [ginkgo](https://github.com/onsi/ginkgo) and 
[gomega](https://github.com/onsi/gomega) - BDD testing
- [gomock](https://github.com/golang/mock) and
[go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - mocking dependencies 
  for testing

Feel free to explore if you are interested in learning any one of these 
libraries or if you want to build highly decoupled apps but have not found your
best approach yet.

---

### How to run

1. Make sure you have [Go](https://go.dev/dl/) and 
[Buf](https://docs.buf.build/installation) installed.
2. Start Postgres database locally on port `:5432`. <br>
The easiest way is to start Postgres container in Docker. <br>
To do it, install [Docker](https://docs.docker.com/get-docker/) first, then 
   run `docker run -p 5432:5432 --name go-example_db -e POSTGRES_USER=user -e 
POSTGRES_PASSWORD=password postgres` in your terminal (username and password 
   are set to defaults which match the values in `configs/*.env` files).
   To stop the container afterwards, run `docker stop go-example_db`.
3. Clone this repository.
4. Run `make init` command to install required dependencies.
5. Run `go mod tidy` command to install libraries used in this project.
6. Build the project using `make build` command
7. Run compiled binary with `bin/server -dev` command to run it in 
   development mode or change `-dev` flag to `-stage` or `-prod` to run in 
   staging and production environments accordingly (for simplicity, the
   configuration is the same for each environment, feel free to change `*.yaml` 
   and `*.env` files in `config/` directory)

---

### More to come

Currently, this is not final form of this project, here are some things you 
will see in the future: 

- Github CI/CD pipeline
- Containerization using Docker 
- More test coverage (mw, server and pkg tests)

Might be implemented:

- TLS configuration for HTTPS support
- Kubernetes cluster
- NGINX proxy 
