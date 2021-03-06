package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(smtpCmd)
}

var smtpCmd = &cobra.Command{
	Use:   "smtp",
	Short: "Send messages using the SMTP server",
	Long:  `Receive the message data, connect to a SMTP server and send it.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Processing SMTP queue")

	},
}
