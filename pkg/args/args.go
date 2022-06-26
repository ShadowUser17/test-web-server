package args

import (
	"flag"
	"os"
)

type Args struct {
	Address  *string
	Location *string
	SSLmode  *bool
	SSLcert  *string
	SSLkey   *string
}

func (args *Args) Parse() *Args {
	args.Address = flag.String("l", "127.0.0.1:8080", "Set listen address.")
	args.Location = flag.String("e", "/", "Set HTTP endpoint.")
	args.SSLmode = flag.Bool("s", false, "Start HTTPS mode.")
	args.SSLcert = flag.String("c", "localhost/cert.pem", "Set SSL cert.")
	args.SSLkey = flag.String("k", "localhost/cert.key", "Set SSL key.")

	flag.Parse()
	return args
}

func (args *Args) CertIsExist() bool {
	if _, err := os.Stat(*args.SSLcert); err != nil {
		return false
	}

	if _, err := os.Stat(*args.SSLkey); err != nil {
		return false
	}

	return true
}
