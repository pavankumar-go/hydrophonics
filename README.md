Hydrophonics
===========
## A Web based Microservice developed with Golang and  uses Gin-Gonic http framework for routing, GORM for interacting with Postgres DB

![Build](https://github.com/pavankumar-go/hydrophonics/workflows/Build/badge.svg?branch=master)

### Requirements for Running it on Mac OS
* Go lang - `version go1.13.6 darwin/amd64` (https://golang.org/dl/)
* Postgres for Mac - `brew install postgresql`
* Go Package Dependencies 
	* `go get -u github.com/gin-gonic/gin`
	* `go get -u github.com/gogo/protobuf`
	* `go get -u github.com/google/uuid`
	* `go get -u github.com/jinzhu/gorm`
	* `go get -u github.com/spf13/viper`
	* `go get -u gopkg.in/birkirb/loggers.v1`

### Steps to Run 
1. Clone the repo 
2. Create a database in Postgres named `hydrophonics`
3. configure your host, port, password in - [config.toml](hack/config/config.toml)
3. cd to `/your/directory/main.go` and `go run main.go` in your terminal

### Useful links for references
1. https://gorm.io

### Roadmap
- [ ] Implement JWT for authentication 
- [ ] Implement RBAC for authorization
- [ ] Dockerize to run it in a container
- [ ] add more attributes to existing tables
