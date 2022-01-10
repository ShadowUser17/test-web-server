package args

import (
	"flag"
)

type Args struct {
	Address *string
	SSLcert *string
	SSLkey  *string
}

func (self *Args) Parse() {
	self.Address = flag.String("l", "127.0.0.1:8080", "127.0.0.1:8080")
	self.SSLcert = flag.String("c", "", "Set SSL cert.")
	self.SSLkey = flag.String("k", "", "Set SSL key.")

	flag.Parse()
}
