package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "uhura",
		Short:   "Message users for several different channels",
		Version: "0.0.1",
		Long:    `Uhura is a CLI that exchanges messages using several different channels. The goal is to reach people the way they feels fit their needs.`,
	}
)

// Execute the trigger for cobra execution process
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().StringVar(&config.cfgFile, "config", "", "config file (default is $HOME/.uhura.yaml)")

	rootCmd.PersistentFlags().String("license", "CC BY-SA", "returns the open source license")
	rootCmd.PersistentFlags().String("codebase", "https://github.com/JBGoldberg/uhura", "returns the source codebase location")
	rootCmd.PersistentFlags().String("author", "Jim Bruno Goldberg <jbgoldberg@nekutima.eu>", "returns the author data")

}
