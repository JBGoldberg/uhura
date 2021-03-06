package cmd

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {

	sendCmd.PersistentFlags().StringP("author", "a", "Jim Bruno Goldberg <jbgoldberg@nekutima.eu>", "author name for copyright attribution")
	sendCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")

	sendCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", sendCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", sendCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Jim Bruno Goldberg <jbgoldberg@nekutima.eu>")
	viper.SetDefault("license", "CC BY-SA")

	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send smtp|telegram",
	Short: "Send messages on queue",
	Long:  `Reads the messages queue to be send and process it.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Sending messages")

		for _, c := range args {

			switch c {
			case "smtp":
				if err := smtpSend(); err != nil {
					return err
				}
				break

			case "telegram":
				log.Errorf("Telegram communications are not implemmented yet")
				return errors.New("not implemented")

			default:
				log.Errorf("I don't know the %s communication channel", c)
				return errors.New("channel unknown")

			}
		}

		return nil

	},
}
