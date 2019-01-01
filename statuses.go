package twitter

type Status struct {
	CreatedAt          string         `json:"created_at"`
	User               User           `json:"user"`
	Text               string         `json:"text"`
	FullText           *string        `json:"full_text"`
	ID                 uint64         `json:"id"`
	IDStr              string         `json:"id_str"`
	InReplyStatusID    *uint64        `json:"in_reply_to_status_id"`
	InReplyStatusIDStr *string        `json:"in_reply_to_status_id_str"`
	InReplyUserID      *uint64        `json:"in_reply_to_user_id"`
	InReplyUserIDStr   *string        `json:"in_reply_to_user_id_str"`
	Entities           Entity         `json:"entities"`
	ExtendedEntities   ExtendedEntity `json:"extended_entities"`
	QuotedStatus       *Status        `json:"quoted_status"`
	Truncated          *bool          `json:"truncated"`
}

type Entity struct {
	Urls         []EntityURL         `json:"urls"`
	Hashtags     []EntityHashtag     `json:"hashtags"`
	UserMentions []EntiryUserMention `json:"user_mentions"`
	Medias       []EntityURL         `json:"media"`
}

type EntityURL struct {
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
	Indices     [2]int `json:"indices"`
}

type EntityHashtag struct {
	Text    string `json:"text"`
	Indices [2]int `json:"indices"`
}

type EntiryUserMention struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	ID         uint64 `json:"id"`
	IDStr      string `json:"id_str"`
	Indices    [2]int `json:"indices"`
}

type ExtendedEntity struct {
	Medias []TwitterMedia `json:"media"`
}

type TwitterMedia struct {
	MediaURL      string `json:"media_url"`
	MediaURLHttps string `json:"media_url_https"`
}

type ReceivedStatus struct {
	Status
	RetweetedStatus *Status `json:"retweeted_status"`
}

const urlStatusesUpdate = "https://api.twitter.com/1.1/statuses/update.json"

func (twitter *Twitter) StatusesUpdate(userParams map[string]string) ([]ReceivedStatus, error) {
	var result []ReceivedStatus
	err := twitter.post(urlStatusesUpdate, userParams, &result)
	return result, err
}
