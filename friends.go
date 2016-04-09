package twitter

const URL_FRIENDS_IDS = "https://api.twitter.com/1.1/friends/ids.json"

func (twitter *Twitter) FriendsIds(userParams map[string]string) (result FfIds, err error) {
	err = twitter.get(URL_FRIENDS_IDS, userParams, &result)
	return result, err
}
