package twitter

type ListMembers struct {
	PreviousCursor    int    `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
	NextCursor        int    `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	Users             []User `json:"users"`
}

const (
	URL_LISTS_STATUSES = "https://api.twitter.com/1.1/lists/statuses.json"
	URL_LISTS_MEMBERS  = "https://api.twitter.com/1.1/lists/members.json"
)

func (twitter *Twitter) ListsStatuses(userParams map[string]string) (result []ReceivedStatus, err error) {
	err = twitter.get_parse(URL_LISTS_STATUSES, userParams, &result)
	return result, err
}
func (twitter *Twitter) ListsMembers(userParams map[string]string) (result ListMembers, err error) {
	err = twitter.get_parse(URL_LISTS_MEMBERS, userParams, &result)
	return result, err
}
