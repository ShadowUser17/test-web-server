package main

import (
	"os"

	"github.com/ShadowUser17/TestWebServer/pkg/args"
	"github.com/ShadowUser17/TestWebServer/pkg/cert"
	"github.com/ShadowUser17/TestWebServer/pkg/server"
)

func main() {
	var params = new(args.Args).Parse()
	var router = server.GetRouter(*params.Location)
	var srv = server.GetServer(*params.Address, router)
	var err error

	if *params.SSLmode {
		if !params.CertIsExist() {
			if _, _, err := cert.MakeCert("./", "localhost"); err != nil {
				router.Logger.Printf("Error: %v\n", err)
				os.Exit(2)
			}
		}

		router.Logger.Printf("Listen: https://%s%s\n", *params.Address, *params.Location)
		err = srv.ListenAndServeTLS(*params.SSLcert, *params.SSLkey)

	} else {
		router.Logger.Printf("Listen: http://%s%s\n", *params.Address, *params.Location)
		err = srv.ListenAndServe()
	}

	if err != nil {
		router.Logger.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
