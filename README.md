# Open Tree Hole Next

Next Generation of OpenTreeHole Written In Golang

An Anonymous BBS

## Features
To be done

## Usage

### build and run

Before building, you should have go (>=1.22) installed on your machine. If not, please visit [golang.org](https://golang.org) to download the latest version.

```shell
git clone https://github.com/OpenTreeHole/treehole_next.git
cd treehole_next
# install swag and generate docs
go install github.com/swaggo/swag/cmd/swag@latest
swag init --parseDependency --parseDepth 1 # to generate the latest docs, this should be run before compiling
# build for debug
go build -o treehole.exe
# build for release
go build -tags "release" -ldflags "-s -w" -o treehole.exe
# run
./treehole.exe
```

Note: You should check [config file](config/config.go) to see if required environment variables are set.

### test

```shell
export MODE=test
go test -v ./tests/...
```

### benchmark

```shell
export MODE=bench
go test -v -benchmem -cpuprofile=cpu.out -benchtime=1s ./benchmarks/... -bench .
```
For documentation, please open http://localhost:8000/docs after running app
## Badge

[//]: # ([![build]&#40;https://github.com/OpenTreeHole/treehole_next/actions/workflows/master.yaml/badge.svg&#41;]&#40;https://github.com/OpenTreeHole/treehole_next/actions/workflows/master.yaml&#41;)
[//]: # ([![dev build]&#40;https://github.com/OpenTreeHole/treehole_next/actions/workflows/dev.yaml/badge.svg&#41;]&#40;https://github.com/OpenTreeHole/treehole_next/actions/workflows/dev.yaml&#41;)

[![stars](https://img.shields.io/github/stars/OpenTreeHole/treehole_next)](https://github.com/OpenTreeHole/treehole_next/stargazers)
[![issues](https://img.shields.io/github/issues/OpenTreeHole/treehole_next)](https://github.com/OpenTreeHole/treehole_next/issues)
[![pull requests](https://img.shields.io/github/issues-pr/OpenTreeHole/treehole_next)](https://github.com/OpenTreeHole/treehole_next/pulls)

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

### Powered by

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/OpenTreeHole/treehole_next/issues/new) or [Submit PRs](https://github.com/OpenTreeHole/treehole_next/compare).

We are now in rapid development, any contribution would be of great help. 
For the developing roadmap, please visit [this issue](https://github.com/OpenTreeHole/treehole_next/issues/1).

### Contributors

This project exists thanks to all the people who contribute.

<a href="https://github.com/OpenTreeHole/treehole_next/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=OpenTreeHole/treehole_next"  alt="contributors"/>
</a>

## Licence

[![license](https://img.shields.io/github/license/OpenTreeHole/treehole_next)](https://github.com/OpenTreeHole/treehole_next/blob/master/LICENSE)
© OpenTreeHole
