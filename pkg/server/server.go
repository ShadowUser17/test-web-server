package server

import (
	"net/http"
	"time"
)

func GetServer(address string, router *Router) *http.Server {
	return &http.Server{
		Addr:              address,
		Handler:           router,
		ErrorLog:          router.Logger,
		IdleTimeout:       time.Second * 60,
		ReadTimeout:       time.Second * 20,
		WriteTimeout:      time.Second * 20,
		ReadHeaderTimeout: time.Second * 20,
	}
}
