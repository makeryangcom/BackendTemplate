package main

import (
	"Platform/framework"
	"Platform/framework/config"
	"Platform/framework/router"
	"context"
	"fmt"
	"github.com/gookit/color"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	Framework.Init()
}

func main() {

	RouterInit := Router.Init()

	var HttpServer = &http.Server{
		Addr:           fmt.Sprintf(":%d", Config.Get.Service.HttpPort),
		Handler:        RouterInit,
		ReadTimeout:    Config.Get.Service.ReadTimeout,
		WriteTimeout:   Config.Get.Service.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := HttpServer.ListenAndServe(); err != nil {
		}
	}()

	log.Println("[main]", color.Green.Text("server..."))

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := HttpServer.Shutdown(ctx); err != nil {
	}
}
