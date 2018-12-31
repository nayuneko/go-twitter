package twitter

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/mrjones/oauth"
)

// Config Config
type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

type Twitter struct {
	config      Config
	RawBodyData string
}

func New(config Config) *Twitter {
	return &Twitter{config: config}
}

func (twitter *Twitter) getConsumer() *oauth.Consumer {
	consumer := oauth.NewConsumer(
		twitter.config.ConsumerKey,
		twitter.config.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
	return consumer
}

func (twitter *Twitter) get(url string, userParams map[string]string, v interface{}) error {
	consumer := twitter.getConsumer()
	token := &oauth.AccessToken{
		Token:  twitter.config.AccessToken,
		Secret: twitter.config.AccessSecret,
	}
	resp, err := consumer.Get(url, userParams, token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	twitter.RawBodyData, err = parse(resp.Body, v)
	return err
}

func (twitter *Twitter) post(url string, userParams map[string]string, v interface{}) error {
	consumer := twitter.getConsumer()
	token := &oauth.AccessToken{
		Token:  twitter.config.AccessToken,
		Secret: twitter.config.AccessSecret,
	}
	resp, err := consumer.Post(url, userParams, token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	twitter.RawBodyData, err = parse(resp.Body, v)
	return err
}

func parse(r io.Reader, v interface{}) (string, error) {
	bits, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bits, &v)
	if err != nil {
		return string(bits), err
	}
	return string(bits), err
}
