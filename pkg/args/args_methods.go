package args

import (
	"flag"
)

type Args struct {
	Address  *string
	Location *string
	SSLcert  *string
	SSLkey   *string
}

func (args *Args) Parse() {
	args.Address = flag.String("l", "127.0.0.1:8080", "Set listen address.")
	args.Location = flag.String("e", "/", "Set HTTP endpoint.")
	args.SSLcert = flag.String("c", "", "Set SSL cert.")
	args.SSLkey = flag.String("k", "", "Set SSL key.")

	flag.Parse()
}

func (args *Args) IsHttps() bool {
	if (*args.SSLcert != "") && (*args.SSLkey != "") {
		return true
	}

	return false
}
