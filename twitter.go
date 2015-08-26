package twitter

import (
	"encoding/json"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"net/http"
)

type Twitter struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

type User struct {
	Name       string
	ScreenName string `json:"screen_name"`
	Id         uint64 `json:"id"`
	IdStr      string `json:"id_str"`
	IconUrl    string `json:"profile_image_url"`
}

type Entity struct {
	Urls         []EntityUrl         `json:"urls"`
	Hashtags     []EntityHashtag     `json:"hashtags"`
	UserMentions []EntiryUserMention `json:"user_mentions"`
	Medias       []EntityUrl         `json:"media"`
}

type EntityUrl struct {
	ExpandedUrl string `json:"expanded_url"`
	DisplayUrl  string `json:"display_url"`
	Indices     [2]int `json:"indices"`
}

type EntityHashtag struct {
	Text    string `json:"text"`
	Indices [2]int `json:"indices"`
}

type EntiryUserMention struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Id         uint64 `json:"id"`
	IdStr      string `json:"id_str"`
	Indices    [2]int `json:"indices"`
}

type ExtendedEntity struct {
	Medias []TwitterMedia `json:"media"`
}

type TwitterMedia struct {
	MediaUrl string `json:"media_url"`
}

type Status struct {
	CreatedAt          string         `json:"created_at"`
	User               User           `json:"user"`
	Text               string         `json:"text"`
	Id                 uint64         `json:"id"`
	IdStr              string         `json:"id_str"`
	InReplyStatusIdStr string         `json:"in_reply_to_status_id_str"`
	InReplyUserIdStr   string         `json:"in_reply_to_user_id_str"`
	Entities           Entity         `json:"entities"`
	ExtendedEntities   ExtendedEntity `json:"extended_entities"`
}

type ReceivedStatus struct {
	Status
	RetweetedStatus *Status `json:"retweeted_status"`
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

func (twitter *Twitter) ListsStatuses(list_id string) ([]ReceivedStatus, error) {
	userParams := map[string]string{"list_id": list_id}
	resp, err := twitter.get("https://api.twitter.com/1.1/lists/statuses.json", userParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result []ReceivedStatus
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bits, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
