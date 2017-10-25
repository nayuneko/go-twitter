package twitter

type Status struct {
	CreatedAt          string         `json:"created_at"`
	User               User           `json:"user"`
	Text               string         `json:"text"`
	FullText           *string        `json:"full_text"`
	Id                 uint64         `json:"id"`
	IdStr              string         `json:"id_str"`
	InReplyStatusIdStr string         `json:"in_reply_to_status_id_str"`
	InReplyUserIdStr   string         `json:"in_reply_to_user_id_str"`
	Entities           Entity         `json:"entities"`
	ExtendedEntities   ExtendedEntity `json:"extended_entities"`
	QuotedStatus       *Status        `json:"quoted_status"`
	Truncated          *bool          `json:"truncated"`
}

type Entity struct {
	Urls         []EntityUrl         `json:"urls"`
	Hashtags     []EntityHashtag     `json:"hashtags"`
	UserMentions []EntiryUserMention `json:"user_mentions"`
	Medias       []EntityUrl         `json:"media"`
}

type EntityUrl struct {
	Url         string `json:"url"`
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
	MediaUrl      string `json:"media_url"`
	MediaUrlHttps string `json:"media_url_https"`
}

type ReceivedStatus struct {
	Status
	RetweetedStatus *Status `json:"retweeted_status"`
}

const URL_STATUSES_UPDATE = "https://api.twitter.com/1.1/statuses/update.json"

func (twitter *Twitter) StatusesUpdate(userParams map[string]string) (result []ReceivedStatus, err error) {
	err = twitter.post(URL_STATUSES_UPDATE, userParams, &result)
	return result, err
}
