# Go port storage service

This is a playground for a simple Go application following DDD/Clean arch principles.

# Domain requirements

- Given a file with ports data [ports.json](testdata/ports.json), the port storage service either creates a new record in a database, or updates the existing one.
  - The file is of unknown size
  - The service has limited resources available
- The app can be build using the provided [Dockerfile](Dockerfile)
- The service supports graceful shutdown

# Additional stuff

- There is an option to run mongoDB in a docker container
- There is a [Docker-compose](docker-compose.yaml) file provided for convenience
- There is a [Makefile](Makefile) provided for convenience

# Usage

## Prerequisites

- go >= 1.18
- docker, docker-compose

## Make targets

### Building locally

  `make build`

#### Testing

`make lint`

`make test`

#### Code formatting

`make fmt`

### Building & running in docker

#### In-memory db

  `make run-inmem`

#### MongoDB

  `make run-mongo`

