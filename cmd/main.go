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

	var srv = server.GetServer(*param.Address)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
