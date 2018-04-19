# Gonitor with WebSocket in Go
[![Build Status](https://travis-ci.org/gonitor/gonitor-websocket.svg?branch=master)](https://travis-ci.org/gonitor/gonitor-websocket)
[![codecov](https://codecov.io/gh/gonitor/gonitor-websocket/branch/master/graph/badge.svg)](https://codecov.io/gh/gonitor/gonitor-websocket)
[![Go Report Card](https://goreportcard.com/badge/github.com/gonitor/gonitor-websocket)](https://goreportcard.com/report/github.com/gonitor/gonitor-websocket)
[![GoDoc](https://godoc.org/github.com/gonitor/gonitor-websocket?status.svg)](https://godoc.org/github.com/gonitor/gonitor-websocket)

Gonitor is fast server monitor service that make monitoring servers easy and simple. Gonitor provides many cool features for monitoring server CPU, Memory, GPU, Disk, Load, Network, etc. If you like or use Gonitor, please star or share it ! 

## Features
- Provides server CPU, Memory, GPU, Disk, Load, Network, Host and Runtime information
- Supports REST API and Swagger UI
- Supports real-time monitoring with Stream or WebSocket
- Supports mutiple-platform build (darwin, windows and linux)
- Supports Docker build and Travis CI

## Installation

Get the gonitor repository

```
go get github.com/gonitor/gonitor-websocket

cd echo $GOPATH/src/github.com/gonitor/gonitor-websocket
```

And install dependencies

```
go get -u github.com/golang/dep/cmd/dep

dep ensure
```

## Running the tests

Run all tests

```
go test ./...
```

Or run all tests with coverage

```
bash coverage.sh
```

## Build and Run

Run main.go
``` bash
go run main.go
# serve at localhost:9000
# serve WebSocket API at localhost:9000/websocket/v1
```

Build and run native binary

```
bash Build.sh

./gonitorWebsocket
```
Build native binary for multiple platforms (darwin, windows and linux)

```
bash BuildMulti.sh
```
## Environment variables

```bash
    # enable production mode, default is false
    env GONI_PRODMODE=true
    
    # disable REST API, default is true
    env GONI_REST=false

    # disable Stream API, default is true
    env GONI_STREAM=false

    # combine three environment variables
    env GONI_PRODMODE=true GONI_REST=false GONI_STREAM=false

    # enable compression, default is true
    env GONI_COMPRESS=false

    # set read buffer size, default equals to 1024
    env GONI_RBSIZE=2048

    # set write buffer size, default equals to 1024
    env GONI_WBSIZE=2048
```
## Docker support 

Build docker image

```
docker build -t gonitor/gonitor-websocket .
```

Run docker container

```
docker run -d --name gonitor-websocket -p 9000:9000 gonitor/gonitor-websocket
```
## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/gonitor/gonitor-websocket/tags). 

## Authors

* **Ai Nguyen** - [AiNguyenKaka](https://github.com/ainguyenkaka)

See also the list of [contributors](https://github.com/gonitor/gonitor-websocket/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

