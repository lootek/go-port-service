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

# Disclaimer

At the time of writing I'm honestly quite unhappy with the result. I wish I had more time and focus to this but life is life.
I must admit I just had to hard stop to get back to my regular work and did't even run that stuff (shame on me).
Leaving aside I'm not used to the rare occasion of creating a software service totally from scratch it was a pretty nice brain exercise. 
I love how the code looks like when it's nicely organized according to SOLID, DDD, clean/hex arch principles, 
yet I was not fortunate enough to be able to work on such code at work since we do have lots of legacy Go codebases 
that not really adhere to those.

What I'd do next if I could:
* try actually running this
  * locally
  * using docker/docker-compose
  * try out the REST server via uploading ports,json file couple times (the same and with some slight changes)
  * try running against mongodb (that needs some actual implementation for the mongo-based repository)
* add more tests
  * for core/ maybe use goconvey to make the tests more readable for non-technical people
  * tests for update logic in repository
  * tests for resilience to malicious input (fuzz tests would be handy here)
* add input validation to the service implementation in core/
