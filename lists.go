package twitter

// ListMembers ListMembers構造体
type ListMembers struct {
	PreviousCursor    int    `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
	NextCursor        int    `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	Users             []User `json:"users"`
}

const (
	urlListsStatuses = "https://api.twitter.com/1.1/lists/statuses.json"
	urlListsMembers  = "https://api.twitter.com/1.1/lists/members.json"
)

// ListsStatuses リスト取得
func (twitter *Twitter) ListsStatuses(userParams map[string]string) ([]ReceivedStatus, error) {
	var result []ReceivedStatus
	err := twitter.get(urlListsStatuses, userParams, &result)
	return result, err
}

// ListsMembers リストメンバー取得
func (twitter *Twitter) ListsMembers(userParams map[string]string) (ListMembers, error) {
	var result ListMembers
	err := twitter.get(urlListsMembers, userParams, &result)
	return result, err
}
