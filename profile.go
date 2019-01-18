package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/viper"
)

var subv *viper.Viper
var logger *log.Logger

func initProfile() error {
	logger = log.New(ioutil.Discard, "", 0)
	if viper.GetBool(KeyVerbose) {
		logger = log.New(os.Stderr, "", log.LstdFlags|log.LUTC)
	}

	profile := viper.GetString(KeyProfile)
	subv = viper.Sub("profiles." + profile)
	if subv == nil {
		return fmt.Errorf("no found for `%v`. give a name with --profile\n", profile)
	}

	return nil
}
