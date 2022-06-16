package server

import (
	"log"
	"net/http"
	"os"
	"sync"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, *log.Logger)

type Router struct {
	mutex    *sync.Mutex
	Logger   *log.Logger
	Handler  HandlerFunc
	Endpoint string
}

func GetRouter(endpoint string) *Router {
	return &Router{
		mutex:    new(sync.Mutex),
		Logger:   log.New(os.Stdout, "", log.Ldate|log.Ltime),
		Handler:  RootHandler,
		Endpoint: endpoint,
	}
}

func (router *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	router.Logger.Printf("%s: %s \"%s\"", req.Method, req.URL, req.UserAgent())

	router.mutex.Lock()
	defer router.mutex.Unlock()

	if router.Endpoint != req.URL.EscapedPath() {
		rw.WriteHeader(http.StatusNotFound)

	} else {
		router.Handler(rw, req, router.Logger)
	}
}
