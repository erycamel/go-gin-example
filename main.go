package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"syscall"

	"github.com/fvbock/endless"

	"github.com/erycamel/go-gin-example/models"
	"github.com/erycamel/go-gin-example/pkg/gredis"
	"github.com/erycamel/go-gin-example/pkg/logging"
	"github.com/erycamel/go-gin-example/pkg/setting"
	"github.com/erycamel/go-gin-example/routers"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/erycamel/go-gin-example

// @license.name MIT
// @license.url https://github.com/erycamel/go-gin-example/blob/master/LICENSE
func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	if runtime.GOOS == "windows" {
		server := &http.Server{
			Addr:           endPoint,
			Handler:        routersInit,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			MaxHeaderBytes: maxHeaderBytes,
		}

		server.ListenAndServe()
		return
	}

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, routersInit)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
