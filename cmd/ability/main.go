package main

import (
	"gitlab.milobella.com/milobella/ability-sdk-go/pkg/ability"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var additionalConfigPath string

//TODO: try to put some common stuff into a separate repository
func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	additionalConfigPath = os.Getenv("ADDITIONAL_CONFIG_PATH")
	if len(additionalConfigPath) != 0 {
		viper.AddConfigPath(additionalConfigPath)
	}

	viper.AddConfigPath(".")
	viper.SetDefault("server.log-level", "info")

	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// TODO: read it in the config when move to viper
	logrus.SetLevel(logrus.DebugLevel)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		logrus.Errorf("Fatal error config file: %s \n", err)
	}

	if level, err := logrus.ParseLevel(viper.GetString("server.log-level")); err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.Warn("Failed to parse the log level. Keeping the logrus default level.")
	}

	logrus.Debugf("Configuration -> %+v", viper.AllSettings())
}

// fun main()
func main() {

	// Initialize server
	server := ability.NewServer("Weather", viper.GetInt("server.port"))
	server.RegisterIntentRule("GET_WEATHER", DefaultIntentHandler)
	server.Serve()
}

func DefaultIntentHandler(_ *ability.Request, resp *ability.Response) {
	// Build the NLG answer
	resp.Nlg.Sentence = "I don't know how to deal with weather for now."
}
