package twitter

import (
	"github.com/mrjones/oauth"
	"net/http"
)

type Twitter struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
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

func (twitter *Twitter) get(url string, userParams map[string]string) (resp *http.Response, err error) {
	consumer := twitter.getConsumer()
	token := &oauth.AccessToken{
		Token:  twitter.AccessToken,
		Secret: twitter.AccessSecret,
	}
	return consumer.Get(url, userParams, token)
}

func (twitter *Twitter) get_parse(url string, userParams map[string]string, v interface{}) error {
	resp, err := twitter.get(url, userParams)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return Utils.Parse(resp.Body, v)
}
