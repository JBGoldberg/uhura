package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Retreives informations about Uhura",
	Long:  `Returns legal and convenience information about Uhura like author, license.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		for _, info := range args {

			switch info {
			case "author":
				fmt.Println("Jim Bruno Goldberg <jbgoldberg@nekutima.eu>")
				break

			case "codebase":
				fmt.Println("https://github.com/JBGoldberg/uhura")
				break

			case "license":
				fmt.Println("CC BY-SA")
				break

			case "version":
				fmt.Println("0.0.1")
				break

			default:
				return errors.New("Requested information is not availiable")

			}
		}

		return nil
	},
}
