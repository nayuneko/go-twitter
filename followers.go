package twitter

const URL_FOLLOWERS_IDS = "https://api.twitter.com/1.1/followers/ids.json"

func (twitter *Twitter) FollowersIds(userParams map[string]string) (result FfIds, err error) {
	err = twitter.get(URL_FOLLOWERS_IDS, userParams, &result)
	return result, err
}
