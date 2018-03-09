package main

import (
	"os"

	"github.com/micro/micro/cmd"

	_ "github.com/micro/go-plugins/client/grpc"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/server/grpc"
)

func init() {
	os.Setenv("MICRO_CLIENT", "grpc")
	os.Setenv("MICRO_SERVER", "grpc")
	os.Setenv("MICRO_REGISTRY", "nats")
}

func main() {
	cmd.Init()
}
