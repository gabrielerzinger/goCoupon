package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/gabrielerzinger/goCoupon/app"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts goCoupon",
	Long:  "starts goCoupon",
	Run: func(cmd *cobra.Command, args []string) {
		var log = logrus.New()

		cmdL := log.WithFields(logrus.Fields{
			"source":    "startCmd",
			"operation": "Run",
		})

		cmdL.Info("starting the app")

		app, err := app.NewApp("0.0.0.0", 8000, config, cmdL)

		cmdL.Info("app created")

		if err != nil {
			cmdL.Fatal(err)
		}

		app.Init()
		cmdL.Info("app started")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
