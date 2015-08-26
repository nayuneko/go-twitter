package twitter

type User struct {
	Name       string
	ScreenName string `json:"screen_name"`
	Id         uint64 `json:"id"`
	IdStr      string `json:"id_str"`
	IconUrl    string `json:"profile_image_url"`
}
