package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"school-walker/walker/pkg/setting"
	"school-walker/walker/routers"
)

func main() {
	log.Printf("pid:%v", os.Getpid())
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
