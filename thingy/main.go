package main

import (
	"../../../tmp/demo/thingy"
	"../../../tmp/demo/thingy/server"
	"../../../tmp/demo/thingy/thingy/cmd"
	"github.com/lileio/lile"
	"github.com/lileio/lile/fromenv"
	"github.com/lileio/pubsub"
	"google.golang.org/grpc"
)

func main() {
	s := &server.ThingyServer{}

	lile.Name("thingy")
	lile.Server(func(g *grpc.Server) {
		thingy.RegisterThingyServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
	})

	cmd.Execute()
}
