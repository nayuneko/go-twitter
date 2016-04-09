package twitter

import (
	"errors"
	"github.com/mrjones/oauth"
	"net/http"
)

type Twitter struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
	RawBodyData    string
}

func NewTwitter(ConsumerKey, ConsumerSecret, AccessToken, AccessSecret string) *Twitter {
	return &Twitter{
		ConsumerKey:    ConsumerKey,
		ConsumerSecret: ConsumerSecret,
		AccessToken:    AccessToken,
		AccessSecret:   AccessSecret,
	}
}
func (twitter *Twitter) getConsumer() *oauth.Consumer {
	consumer := oauth.NewConsumer(
		twitter.ConsumerKey,
		twitter.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
	return consumer
}

const (
	METHOD_GET  = "GET"
	METHOD_POST = "POST"
)

func (twitter *Twitter) request(method string, url string, userParams map[string]string, v interface{}) error {
	consumer := twitter.getConsumer()
	token := &oauth.AccessToken{
		Token:  twitter.AccessToken,
		Secret: twitter.AccessSecret,
	}
	var err error = nil
	var resp *http.Response = nil
	if method == METHOD_GET {
		resp, err = consumer.Get(url, userParams, token)
	} else if method == METHOD_POST {
		resp, err = consumer.Post(url, userParams, token)
	} else {
		return errors.New("Unknown method")
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	twitter.RawBodyData, err = Utils.Parse(resp.Body, v)
	return err
}
func (twitter *Twitter) get(url string, userParams map[string]string, v interface{}) error {
	return twitter.request(METHOD_GET, url, userParams, v)
}
func (twitter *Twitter) post(url string, userParams map[string]string, v interface{}) error {
	return twitter.request(METHOD_POST, url, userParams, v)
}
