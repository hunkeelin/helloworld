# helloworld
[![Go Report Card](https://goreportcard.com/badge/github.com/hunkeelin/helloworld)](https://goreportcard.com/report/github.com/hunkeelin/helloworld)
[![GoDoc](https://godoc.org/github.com/hunkeelin/helloworld?status.svg)](https://godoc.org/github.com/hunkeelin/helloworld)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hunkeelin/helloworld/master/LICENSE)
[![CircleCI](https://circleci.com/gh/hunkeelin/helloworld.svg?style=shield)](https://circleci.com/gh/hunkeelin/helloworld)


Simple webapp that listens on port `8080` and prints `hello world` and timestamp

### Build and Run
```bash
go build -o helloworld helloworld.go
./helloworld
```

### Helm
```bash
helm upgrade -i helloworld . -n ${namespace}
```
