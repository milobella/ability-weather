package main

import (
	"gitlab.milobella.com/milobella/ability-sdk-go/pkg/ability"
	"gitlab.milobella.com/milobella/weather-ability/pkg/tools/weather"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var additionalConfigPath string
var weatherClient *weather.Client
var weatherSentencesPresent map[string]string
var weatherSentencesFuture map[string]string

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

	weatherClient = weather.NewClient(
		viper.GetString("tools.weather.host"),
		viper.GetInt("tools.weather.port"))

	weatherSentencesPresent := make(map[string]string)
	weatherSentencesPresent["thunderstorm with light rain"] = "There is thunderstorm with light rain."
	weatherSentencesPresent["thunderstorm with rain"] = "There is thunderstorm with rain."
	weatherSentencesPresent["thunderstorm with heavy rain"] = "There is thunderstorm with heavy rain."
	weatherSentencesPresent["light thunderstorm"] = "There is light thunderstorm."
	weatherSentencesPresent["thunderstorm"] = "There is thunderstorm."
	weatherSentencesPresent["heavy thunderstorm"] = "There is heavy thunderstorm."
	weatherSentencesPresent["ragged thunderstorm"] = "There is ragged thunderstorm."
	weatherSentencesPresent["thunderstorm with light drizzle"] = "There is thunderstorm with light drizzle."
	weatherSentencesPresent["thunderstorm with drizzle"] = "There is thunderstorm with drizzle."
	weatherSentencesPresent["thunderstorm with heavy drizzle"] = "There is thunderstorm with heavy drizzle."
	weatherSentencesPresent["light intensity drizzle"] = "There is light intensity drizzle."
	weatherSentencesPresent["drizzle"] = "There is drizzle."
	weatherSentencesPresent["heavy intensity drizzle"] = "There is heavy intensity drizzle."
	weatherSentencesPresent["light intensity drizzle rain"] = "There is light intensity drizzle rain."
	weatherSentencesPresent["drizzle rain"] = "There is drizzle rain."
	weatherSentencesPresent["heavy intensity drizzle rain"] = "There is heavy intensity drizzle rain."
	weatherSentencesPresent["shower rain and drizzle"] = "There is shower rain and drizzle."
	weatherSentencesPresent["heavy shower rain and drizzle"] = "There is heavy shower rain and drizzle."
	weatherSentencesPresent["shower drizzle"] = "There is shower drizzle."
	weatherSentencesPresent["light rain"] = "There is light rain."
	weatherSentencesPresent["moderate rain"] = "There is moderate rain."
	weatherSentencesPresent["heavy intensity rain"] = "There is heavy intensity rain."
	weatherSentencesPresent["very heavy rain"] = "There is very heavy rain."
	weatherSentencesPresent["extreme rain"] = "There is extreme rain."
	weatherSentencesPresent["freezing rain"] = "There is freezing rain."
	weatherSentencesPresent["light intensity shower rain"] = "There is light intensity shower rain."
	weatherSentencesPresent["shower rain"] = "There is shower rain."
	weatherSentencesPresent["heavy intensity shower rain"] = "There is heavy intensity shower rain."
	weatherSentencesPresent["ragged shower rain"] = "There is ragged shower rain."
	weatherSentencesPresent["light snow"] = "There is light snow."
	weatherSentencesPresent["snow"] = "There is snow."
	weatherSentencesPresent["heavy snow"] = "There is heavy snow."
	weatherSentencesPresent["sleet"] = "There is sleet."
	weatherSentencesPresent["light shower sleet"] = "There is light shower sleet."
	weatherSentencesPresent["shower sleet"] = "There is shower sleet."
	weatherSentencesPresent["light rain and snow"] = "There is light rain and snow."
	weatherSentencesPresent["rain and snow"] = "There is rain and snow."
	weatherSentencesPresent["light shower snow"] = "There is light shower snow."
	weatherSentencesPresent["shower snow"] = "There is shower snow."
	weatherSentencesPresent["heavy shower snow"] = "There is heavy shower snow."
	weatherSentencesPresent["mist"] = "There is mist."
	weatherSentencesPresent["smoke"] = "There is smoke."
	weatherSentencesPresent["haze"] = "There is haze."
	weatherSentencesPresent["sand dust whirls"] = "There is sand dust whirls."
	weatherSentencesPresent["fog"] = "There is fog."
	weatherSentencesPresent["sand"] = "There is sand."
	weatherSentencesPresent["dust"] = "There is dust."
	weatherSentencesPresent["volcanic ash"] = "There is volcanic ash."
	weatherSentencesPresent["squalls"] = "There is squalls."
	weatherSentencesPresent["tornado"] = "There is tornado."
	weatherSentencesPresent["clear"] = "There is clear."
	weatherSentencesPresent["few clouds"] = "There is few clouds."
	weatherSentencesPresent["scattered clouds"] = "There is scattered clouds."
	weatherSentencesPresent["broken clouds"] = "There is broken clouds."
	weatherSentencesPresent["overcast clouds"] = "There is overcast clouds."
	weatherSentencesFuture := make(map[string]string)
	weatherSentencesFuture["thunderstorm with light rain"] = "There will be thunderstorm with light rain."
	weatherSentencesFuture["thunderstorm with rain"] = "There will be thunderstorm with rain."
	weatherSentencesFuture["thunderstorm with heavy rain"] = "There will be thunderstorm with heavy rain."
	weatherSentencesFuture["light thunderstorm"] = "There will be light thunderstorm."
	weatherSentencesFuture["thunderstorm"] = "There will be thunderstorm."
	weatherSentencesFuture["heavy thunderstorm"] = "There will be heavy thunderstorm."
	weatherSentencesFuture["ragged thunderstorm"] = "There will be ragged thunderstorm."
	weatherSentencesFuture["thunderstorm with light drizzle"] = "There will be thunderstorm with light drizzle."
	weatherSentencesFuture["thunderstorm with drizzle"] = "There will be thunderstorm with drizzle."
	weatherSentencesFuture["thunderstorm with heavy drizzle"] = "There will be thunderstorm with heavy drizzle."
	weatherSentencesFuture["light intensity drizzle"] = "There will be light intensity drizzle."
	weatherSentencesFuture["drizzle"] = "There will be drizzle."
	weatherSentencesFuture["heavy intensity drizzle"] = "There will be heavy intensity drizzle."
	weatherSentencesFuture["light intensity drizzle rain"] = "There will be light intensity drizzle rain."
	weatherSentencesFuture["drizzle rain"] = "There will be drizzle rain."
	weatherSentencesFuture["heavy intensity drizzle rain"] = "There will be heavy intensity drizzle rain."
	weatherSentencesFuture["shower rain and drizzle"] = "There will be shower rain and drizzle."
	weatherSentencesFuture["heavy shower rain and drizzle"] = "There will be heavy shower rain and drizzle."
	weatherSentencesFuture["shower drizzle"] = "There will be shower drizzle."
	weatherSentencesFuture["light rain"] = "There will be light rain."
	weatherSentencesFuture["moderate rain"] = "There will be moderate rain."
	weatherSentencesFuture["heavy intensity rain"] = "There will be heavy intensity rain."
	weatherSentencesFuture["very heavy rain"] = "There will be very heavy rain."
	weatherSentencesFuture["extreme rain"] = "There will be extreme rain."
	weatherSentencesFuture["freezing rain"] = "There will be freezing rain."
	weatherSentencesFuture["light intensity shower rain"] = "There will be light intensity shower rain."
	weatherSentencesFuture["shower rain"] = "There will be shower rain."
	weatherSentencesFuture["heavy intensity shower rain"] = "There will be heavy intensity shower rain."
	weatherSentencesFuture["ragged shower rain"] = "There will be ragged shower rain."
	weatherSentencesFuture["light snow"] = "There will be light snow."
	weatherSentencesFuture["snow"] = "There will be snow."
	weatherSentencesFuture["heavy snow"] = "There will be heavy snow."
	weatherSentencesFuture["sleet"] = "There will be sleet."
	weatherSentencesFuture["light shower sleet"] = "There will be light shower sleet."
	weatherSentencesFuture["shower sleet"] = "There will be shower sleet."
	weatherSentencesFuture["light rain and snow"] = "There will be light rain and snow."
	weatherSentencesFuture["rain and snow"] = "There will be rain and snow."
	weatherSentencesFuture["light shower snow"] = "There will be light shower snow."
	weatherSentencesFuture["shower snow"] = "There will be shower snow."
	weatherSentencesFuture["heavy shower snow"] = "There will be heavy shower snow."
	weatherSentencesFuture["mist"] = "There will be mist."
	weatherSentencesFuture["smoke"] = "There will be smoke."
	weatherSentencesFuture["haze"] = "There will be haze."
	weatherSentencesFuture["sand dust whirls"] = "There will be sand dust whirls."
	weatherSentencesFuture["fog"] = "There will be fog."
	weatherSentencesFuture["sand"] = "There will be sand."
	weatherSentencesFuture["dust"] = "There will be dust."
	weatherSentencesFuture["volcanic ash"] = "There will be volcanic ash."
	weatherSentencesFuture["squalls"] = "There will be squalls."
	weatherSentencesFuture["tornado"] = "There will be tornado."
	weatherSentencesFuture["clear"] = "There will be clear."
	weatherSentencesFuture["few clouds"] = "There will be few clouds."
	weatherSentencesFuture["scattered clouds"] = "There will be scattered clouds."
	weatherSentencesFuture["broken clouds"] = "There will be broken clouds."
	weatherSentencesFuture["overcast clouds"] = "There will be overcast clouds."

}

// fun main()
func main() {
	// Initialize server
	server := ability.NewServer("Weather", viper.GetInt("server.port"))
	server.RegisterIntentRule("GET_WEATHER", DefaultIntentHandler)
	server.Serve()
}

func DefaultIntentHandler(_ *ability.Request, resp *ability.Response) {
	// TODO: take this parameter from user preferences or from request
	city := "Cannes"

	// Get weather calling the weather api
	data, err := weatherClient.GetWeather(city)
	if err != nil {
		resp.Nlg.Sentence = "An error occurred retrieving the weather for {{city}}."
		resp.Nlg.Params = []ability.NLGParam{{
			Name: "city",
			Value: city,
			Type: "string",
		}}
		return
	}

	// Retrieve the weather sentence.
	var weatherSentence string
	if val, ok := weatherSentencesPresent[data.Weather]; ok {
		weatherSentence = val
	} else {
		weatherSentence = weatherSentencesPresent["clear"]
	}

	// Build the NLG answer
	resp.Nlg.Sentence = "In {{city}} now, the temperature is {{temperature}}. {{weather_sentence}}"
	resp.Nlg.Params = []ability.NLGParam{{
		Name: "city",
		Value: city,
		Type: "string",
	},{
		Name: "temperature",
		Value: data.Temperature,
		Type: "string",
	},{
		Name: "weather_sentence",
		Value: weatherSentence,
		Type: "inner",
	}}
}
