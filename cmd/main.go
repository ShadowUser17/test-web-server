package main

import (
	"fmt"
	"os"

	"github.com/ShadowUser17/TestWebServer/pkg/args"
	"github.com/ShadowUser17/TestWebServer/pkg/server"
)

func main() {
	var param = args.Args{}
	param.Parse()

	var srv = server.GetServer(*param.Address, *param.Location)
	var err error

	if param.IsHttps() {
		err = srv.ListenAndServeTLS(*param.SSLcert, *param.SSLkey)

	} else {
		err = srv.ListenAndServe()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
