# Example go project

___

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
- [mockery](https://github.com/vektra/mockery) and
[go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - mocking dependencies 
  for testing

Feel free to explore if you are interested in learning any one of these 
libraries or if you want to build highly decoupled apps but have not found your
best approach yet.

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
