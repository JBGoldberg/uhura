package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send messages on queue",
	Long:  `Reads the messages queue to be send and process it.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Processing SMTP queue")

	},
}
