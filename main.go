package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexflint/go-arg"
)

type Args struct {
	Dir  string `arg:"-d" default:"."         help:"directory to serve"`
	Host string `arg:"-a" default:"127.0.0.1" help:"host to bind to"`
	Port string `arg:"-p" default:"5000"      help:"port to serve on"`
}

func main() {
	var args Args
	arg.MustParse(&args)

	var addr = fmt.Sprintf("%s:%s", args.Host, args.Port)

	log.Printf("Serving %s on http://%s\n", args.Dir, addr)

	log.Fatal(
		http.ListenAndServe(
			addr,
			http.FileServer(http.Dir(args.Dir)),
		),
	)
}
