package server

import (
	"net/http"
	"strings"
)

type ApiHandler struct {
}

func (self *ApiHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	switch {
	case req.Method == http.MethodGet && strings.EqualFold(req.URL.Path, "/"):
		self.GetRoot(rw, req)
		return

	case req.Method == http.MethodGet && strings.EqualFold(req.URL.Path, "/test"):
		self.GetTest(rw, req)
		return

		//case req.Method == http.MethodGet && strings.EqualFold(req.URL.Path, "/api"):
	}
}

func (self *ApiHandler) GetRoot(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/api", http.StatusMovedPermanently)
}

func (self *ApiHandler) GetTest(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Server working!"))
}
