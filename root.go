package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var flagOffset *string
var flagTopic *string
var flagPartition *int32
var flagGroup *string
var flagDataCert *string
var flagCACert *string
var flagWithoutGroup *bool

const (
	KeyVerbose     = "verbose"
	KeyProfile     = "profile"
	KeyTLSInsecure = "tls-insecure"
	KeyHosts       = "hosts"
	KeyTopic       = "topic"
	KeyOffset      = "offset"
	KeyPartition   = "partition"
	KeyGroup       = "group"
	KeyDataCert    = "data-cert"
	KeyCACert      = "ca-cert"
)

var RootCmd = &cobra.Command{
	Use:   "kafka-cli",
	Short: "kafka utility to consume",
}

func init() {
	cobra.OnInitialize(initConfig)
	pflags := RootCmd.PersistentFlags()

	pflags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/kafka-cli.yaml)")

	pflags.Bool("verbose", false, `Logging verbosely`)
	pflags.String("profile", "", `Profile name`)
	pflags.Bool("tls-insecure", false, `Allow insecure SSL connection`)

	viper.BindPFlag(KeyVerbose, RootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag(KeyProfile, pflags.Lookup("profile"))
	viper.BindPFlag(KeyTLSInsecure, pflags.Lookup("tls-insecur"))

	//
	flagTopic = pflags.String("topic", "", `Topic name`)
	flagPartition = pflags.Int32("partition", 0, `Partition number`)
	flagGroup = pflags.String("group", "", `Consumer group name`)
	flagOffset = pflags.String("offset", "newest", "`newest`, `oldest` or positive numer")

	//
	flagDataCert = pflags.String("data-cert-path", "", `Data cert path`)
	flagCACert = pflags.String("ca-cert-path", "", `CA cert path`)
}

func main() {
	Execute()
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("kafka-cli")
		viper.AddConfigPath("$HOME/.config")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
