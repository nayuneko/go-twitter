package twitter

type FfIds struct {
	Ids               []uint64 `json:"ids"`
	NextCursor        uint64   `json:"next_cursor"`
	NextCursorStr     string   `json:"next_cursor_str"`
	PreviousCursor    uint64   `json:"previous_cursor"`
	PreviousCursorStr string   `json:"previous_cursor_str"`
}

const (
	urlFollowersIds = "https://api.twitter.com/1.1/followers/ids.json"
	urlFriendsIds   = "https://api.twitter.com/1.1/friends/ids.json"
)

// FollowersIds フォロワーIDリスト取得
func (twitter *Twitter) FollowersIds(userParams map[string]string) (FfIds, error) {
	var result FfIds
	err := twitter.get(urlFollowersIds, userParams, &result)
	return result, err
}

// FriendsIds フォローIDリスト取得
func (twitter *Twitter) FriendsIds(userParams map[string]string) (FfIds, error) {
	var result FfIds
	err := twitter.get(urlFriendsIds, userParams, &result)
	return result, err
}
