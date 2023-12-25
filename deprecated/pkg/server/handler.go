package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func RootHandler(rw http.ResponseWriter, req *http.Request, logger *log.Logger) {
	rw.Header().Add("Content-Type", "text/html")
	rw.WriteHeader(http.StatusOK)

	switch req.Method {
	case http.MethodGet:
		GetDeleteHandler(rw, req, logger)
		return

	case http.MethodPut:
		PutPostHandler(rw, req, logger)
		return

	case http.MethodHead:
		return

	case http.MethodPost:
		PutPostHandler(rw, req, logger)
		return

	case http.MethodDelete:
		GetDeleteHandler(rw, req, logger)
		return

	default:
		var message = fmt.Sprintf("Error: not supported method: %s\n", req.Method)
		http.Error(rw, message, http.StatusMethodNotAllowed)
		logger.Println(message)
	}
}

func PutPostHandler(rw http.ResponseWriter, req *http.Request, logger *log.Logger) {
	fmt.Fprintf(rw, "%s: %s\nBODY: ", req.Method, req.URL.Path)
	io.Copy(rw, req.Body)
	rw.Write([]byte("\n"))
}

func GetDeleteHandler(rw http.ResponseWriter, req *http.Request, logger *log.Logger) {
	fmt.Fprintf(rw, "%s: %s\n", req.Method, req.URL.Path)
}
