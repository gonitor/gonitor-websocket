# Fast Server Monitor with WebSocket in Go
[![Build Status](https://travis-ci.org/gonitor/gonitor-websocket.svg?branch=master)](https://travis-ci.org/gonitor/gonitor-websocket)
[![codecov](https://codecov.io/gh/gonitor/gonitor-websocket/branch/master/graph/badge.svg)](https://codecov.io/gh/gonitor/gonitor-websocket)
[![Go Report Card](https://goreportcard.com/badge/github.com/gonitor/gonitor-websocket)](https://goreportcard.com/report/github.com/gonitor/gonitor-websocket)
[![GoDoc](https://godoc.org/github.com/gonitor/gonitor-websocket?status.svg)](https://godoc.org/github.com/gonitor/gonitor-websocket)

Gonitor is fast server monitor service that make monitoring servers easy and simple. Gonitor provides many cool features for monitoring server CPU, Memory, GPU, Disk, Load, Network, etc. If you like or use Gonitor, please star or share it ! 

## Features
- Provides server CPU, Memory, GPU, Disk, Load, Network, Host and Runtime information
- Supports real-time monitoring with WebSocket
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
sh coverage.sh
```

## Build and Run

Run main.go
``` bash
go run main.go
# serve at localhost:9000
# serve WebSocket API at localhost:9000/websocket
```

Build and run native binary

```
sh Build.sh

./gonitor-websocket
```
Build native binary for multiple platforms (darwin, windows and linux)

```
sh BuildMulti.sh
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

