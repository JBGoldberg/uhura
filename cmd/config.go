package cmd

import (
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type configSMTP struct {
	serverHost string
	serverPort uint
	clientHost string
}

type configAMPQQueues struct {
	smtp     string
	telegram string
}

type configAMPQ struct {
	username        string
	password        string
	serverHost      string
	serverPort      int
	serverAdminPort int
	queues          configAMPQQueues
}

type configuration struct {
	cfgFile string
	smtp    configSMTP
	ampq    configAMPQ
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
		ampq: configAMPQ{
			username:        "tonnystark",
			password:        "pepperssmell",
			serverHost:      "ampq.somewhere.com",
			serverPort:      5672,
			serverAdminPort: 15672,
			queues: configAMPQQueues{
				smtp:     "email-send",
				telegram: "telegram-send",
			},
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
		log.Debugf("Using config file:", viper.ConfigFileUsed())

		config.smtp.clientHost = viper.GetString("smtp.client-host")

		config.smtp.serverHost = viper.GetString("smtp.server-host")
		config.smtp.serverPort = viper.GetUint("smtp.server-port")

		config.ampq.username = viper.GetString("ampq.username")
		config.ampq.password = viper.GetString("ampq.password")
		config.ampq.serverHost = viper.GetString("ampq.server-host")
		config.ampq.serverPort = viper.GetInt("ampq.server-port")
		config.ampq.serverAdminPort = viper.GetInt("ampq.server-admin-port")

	} else {
		log.Error(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
