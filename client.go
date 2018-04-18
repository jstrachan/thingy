package thingy

import (
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	cm     = &sync.Mutex{}
	Client ThingyClient
)

func GetThingyClient() (ThingyClient, error) {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client, nil
	}

	logrus.Info("Creating thingy gRPC client")
	conn, err := grpc.Dial("thingy:80", grpc.DialOption(grpc.WithInsecure()))
	if err != nil {
		return nil, err
	}

	cli := NewThingyClient(conn)
	Client = cli
	return cli, nil
}
