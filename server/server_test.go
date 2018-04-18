package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"../../../tmp/demo/thingy"
	"github.com/lileio/lile"
)

var s = ThingyServer{}
var cli thingy.ThingyClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		thingy.RegisterThingyServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = thingy.NewThingyClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
