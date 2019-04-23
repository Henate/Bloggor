package main

import (
	"fmt"
	"github.com/Henate/Bloggor/pkg/setting"
	"github.com/Henate/Bloggor/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//heartbeat.CronInit()
	s.ListenAndServe()
}