package main

import (
	"net/http"
)

var Server = http.Server{
	Addr:    addr + ":80",
	Handler: Handler(),
}

var ServerTLS = http.Server{
	Addr:    addr + ":443",
	Handler: HandlerTLS(),
}
