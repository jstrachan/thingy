package cmd

import (
	"../../../tmp/demo/thingy/subscribers"
	"github.com/lileio/lile"
	"github.com/lileio/pubsub"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up runs both RPC and pubub subscribers",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			logrus.Fatal(lile.Serve())

		}()

		pubsub.Subscribe(&subscribers.ThingyServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
}
