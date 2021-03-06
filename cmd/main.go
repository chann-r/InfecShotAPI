package main

import (
	"InfecShotAPI/pkg/server"
	"flag"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":80", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Serve(addr)
}
