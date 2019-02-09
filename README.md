# go-skeleton

> Skeleton for Golang projects complete with database layer, gRPC, logging, and more.

[![CircleCI](https://badgen.net/circleci/github/mandocaesar/go-skeleton)](https://circleci.com/gh/mandocaesar/go-skeleton)

## What's in the project

- [ ] gRPC
- [x] Rest API
- [ ] Messaging (Rabbit MQ)
- [x] Data base layer - GORM
- [ ] Data base layer - Repository
- [x] Data base layer - Factory
- [x] Dependency Injection
- [x] Generic Logging - Logrush
- [ ] Generic Logging - fan out multiple channel
- [ ] APM integration
- [ ] Example Project

## Prerequisites

- [Go](https://golang.org/) (1.11.5+)
- [dep](https://golang.github.io/dep/) (0.5.0+)

## Getting started

To use this started, clone this repository and rename the project directory to your liking.

```bash
$ git clone https://github.com/mandocaesar/go-skeleton.git <your-project-name>
```

Next, `cd` into the project folder and install the required dependencies.

```bash
go get

# ...or, if you use dep:
dep ensure
```

You're ready to go!

## Commands

```bash
# builds a production copy of the project
go build

# TODO: add more commands
```

## License

[MIT](LICENSE).