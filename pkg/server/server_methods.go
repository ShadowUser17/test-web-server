package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func GetServer(address string, location string) *http.Server {
	var mux = http.NewServeMux()
	mux.HandleFunc(location, ApiHandler)

	var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &http.Server{
		Addr:              address,
		Handler:           mux,
		ErrorLog:          logger,
		IdleTimeout:       time.Second * 60,
		ReadTimeout:       time.Second * 20,
		WriteTimeout:      time.Second * 20,
		ReadHeaderTimeout: time.Second * 20,
	}
}

func ApiHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/html")

	switch {
	case req.Method == http.MethodGet:
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "<h2>%s: %s</h2>\n", req.Method, req.RequestURI)
		fmt.Fprintf(rw, "<h2>Headers:</h2>\n")
		GetHeaders(rw, req)
		return

	case req.Method == http.MethodPost:
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "<h2>%s: %s</h2>\n", req.Method, req.RequestURI)
		fmt.Fprintf(rw, "<h2>Headers:</h2>\n")
		GetHeaders(rw, req)

		fmt.Fprintf(rw, "<h2>Body:</h2>\n")
		io.Copy(rw, req.Body)
		return

	case req.Method == http.MethodPut:
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "<h2>%s: %s</h2>\n", req.Method, req.RequestURI)
		fmt.Fprintf(rw, "<h2>Headers:</h2>\n")
		GetHeaders(rw, req)

		fmt.Fprintf(rw, "<h2>Body:</h2>\n")
		io.Copy(rw, req.Body)
		return

	case req.Method == http.MethodDelete:
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "<h2>%s: %s</h2>\n", req.Method, req.RequestURI)
		fmt.Fprintf(rw, "<h2>Headers:</h2>\n")
		GetHeaders(rw, req)
		return

	case req.Method == http.MethodHead:
		rw.WriteHeader(http.StatusOK)
		return

	default:
		http.Error(rw, fmt.Sprintf("Not supported method: %s\n", req.Method), http.StatusMethodNotAllowed)
	}
}

func GetHeaders(rw http.ResponseWriter, req *http.Request) {
	for key := range req.Header {
		fmt.Fprintf(rw, "<br>%s: %s\n", key, req.Header.Get(key))
	}
}
