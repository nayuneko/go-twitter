package twitter

type User struct {
	Name            string
	ScreenName      string `json:"screen_name"`
	ID              uint64 `json:"id"`
	IDStr           string `json:"id_str"`
	ProfileImageURL string `json:"profile_image_url"`
}
