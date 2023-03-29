# Getting Started
### Setup Development Environment
* Goland: https://www.jetbrains.com/go/
* Go 1.16: https://golang.org/doc/install
* Set GOPATH: https://golang.org/doc/gopath_code

    GOPATH = $PROJECT_PATH/go
### Framework
* Gin-gonic: https://github.com/gin-gonic/gin
### How To Run
* Fetch go modules:
```bash
go get
```
* Run server

    Start Go instances as usual
* Generate Swagger UI
```bash
go get -u github.com/swaggo/swag/cmd/swag
swag init
```
Swagger UI: /swagger/index.html
