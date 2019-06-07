package weather

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)


type Client struct {
	host    string
	port    int
	url     string
	client  http.Client
}

func NewClient(host string, port int) *Client {
	url := fmt.Sprintf("http://%s:%d", host, port)
	return &Client{host: host, port: port, url: url, client: http.Client{}}
}

func (c Client) makeRequest(method string, endpoint string) (result []byte, err error) {

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		logrus.Error(err)
		return
	}

	resp, err := c.client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}
	output, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logrus.Error(err)
		return
	}

	return output, err
}

func (c Client) GetWeather(city string) (wd WeatherData, err error) {
	endpoint := fmt.Sprintf("%s/weather?q=%s", c.url, city)

	// Proceed the request
	body, err := c.makeRequest("GET", endpoint)

	err = json.Unmarshal(body, &wd)
	return
}

func (c Client) GetForecast(city string) (fc Forecast, err error) {
	endpoint := fmt.Sprintf("%s/forecast?q=%s", c.url, city)

	// Proceed the request
	body, err := c.makeRequest("GET", endpoint)

	err = json.Unmarshal(body, &fc)
	return
}
