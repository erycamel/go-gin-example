.PHONY: build clean tool lint help
GCC=go
GCMD=run
GPATH=main.go
GO_VERSION=1.11.4

run:
	make build
	$(GCC) $(GCMD) $(GPATH)
	
all: build

build:
	@go build -v .

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf go-gin-example
	go clean -i .

install_pkg:
	go get -u github.com/gin-gonic/gin
	go get -u github.com/go-ini/ini
	go get -u github.com/Unknwon/com
	go get -u github.com/astaxie/beego/validation
	go get -u github.com/dgrijalva/jwt-go
	go get -u github.com/gomodule/redigo/redis
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/gin-swagger/swaggerFiles
	go get -u github.com/fvbock/endless
	go get -u github.com/boombuler/barcode
	go get -u github.com/boombuler/barcode/qr
	go get -u golang.org/x/image/math/fixed
	go get -u golang.org/x/image/font

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
