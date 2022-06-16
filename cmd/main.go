package main

import (
	"os"

	"github.com/ShadowUser17/TestWebServer/pkg/args"
	"github.com/ShadowUser17/TestWebServer/pkg/server"
)

func main() {
	var params = new(args.Args).Parse()
	var router = server.GetRouter(*params.Location)
	var srv = server.GetServer(*params.Address, router)
	var err error

	if !params.IsHttps() {
		err = srv.ListenAndServe()

	} else {
		err = srv.ListenAndServeTLS(*params.SSLcert, *params.SSLkey)
	}

	if err != nil {
		router.Logger.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
