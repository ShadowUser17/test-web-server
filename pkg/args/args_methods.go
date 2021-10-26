package args

import (
	"flag"
)

type Args struct {
	Address *string
}

func (self *Args) Parse() {
	self.Address = flag.String("l", "127.0.0.1:8080", "127.0.0.1:8080")

	flag.Parse()
}
