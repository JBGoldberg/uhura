package cmd

import (
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type configSMTP struct {
	serverHost string
	serverPort int
	clientHost string
}

type configuration struct {
	cfgFile string
	smtp    configSMTP
}

// Holds the configuration
var (
	config = configuration{
		cfgFile: "",
		smtp: configSMTP{
			serverHost: "smtp.somewhere.com",
			serverPort: 25,
			clientHost: "thispc.somewhere.com",
		},
	}
)

func initConfig() {
	if config.cfgFile != "" {
		viper.SetConfigFile(config.cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".uhura")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())

		config.smtp.clientHost = viper.GetString("smtp.client-host")

		config.smtp.serverHost = viper.GetString("smtp.server-host")
		config.smtp.serverPort = viper.GetInt("smtp.server-port")

	} else {
		log.Error(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
