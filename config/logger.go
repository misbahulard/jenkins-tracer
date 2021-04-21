package config

import (
	"fmt"
	"os"

	"github.com/misbahulard/jenkins-tracer/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ConfigureLogger() error {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if viper.GetBool("log.debug") {
		log.SetLevel(log.DebugLevel)
	}

	// check if we need to store log to file
	if viper.GetBool("log.file.enable") {
		if viper.GetString("log.file.path") == "" {
			fmt.Println("You enable the file logging, but not define the log path")
			os.Exit(1)
		}

		err := util.CreateDirectoryByFile(viper.GetString("log.file.path"))
		if err != nil {
			return err
		}

		file, err := os.OpenFile(viper.GetString("log.file.path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		log.SetOutput(file)
	}

	log.Info("Logger: ok")
	return nil
}
