package main

import (
	"log"

	"github.com/kyleconroy/hellogrpc"
)

func main() {
	if err := hellogrpc.Serve(); err != nil {
		log.Fatalf("error serving app: %s", err)
	}
}
