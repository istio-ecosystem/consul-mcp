package main

import (
	"flag"

	"github.com/istio-ecosystem/consul-mcp/pkg/server"
)

const (
	defaultListeningAddress = "0.0.0.0:1109"
	defaultConsulAddress    = "127.0.0.1:8500"
)

func main() {
	listeningAddress := flag.String("listeningAddress", defaultListeningAddress, "Listening Address")
	consulAddress := flag.String("consulAddress", defaultConsulAddress, "Consul Address")
	flag.Parse()
	consulMcpServer := server.NewServer(*listeningAddress, *consulAddress)
	if err := consulMcpServer.Start(); err != nil {
		consulMcpServer.Stop()
	}
}
