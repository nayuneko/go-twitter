package twitter

const (
	URL_LISTS_STATUSES = "https://api.twitter.com/1.1/lists/statuses.json"
)

func (twitter *Twitter) ListsStatuses(userParams map[string]string) (result []ReceivedStatus, err error) {
	err = twitter.get_parse(URL_LISTS_STATUSES, userParams, &result)
	return result, err
}
