package anaconda

type TwitterEntities struct {
	Hashtags []struct {
		Indices []int `json:"indices"`
		Text    string `json:"text"`
	} `json:"hashtags"`

	Urls []struct {
		Indices      []int `json:"indices"`
		Url          string `json:"url"`
		Display_url  string `json:"display_url"`
		Expanded_url string `json:"expanded_url"`
	} `json:"urls"`

	User_mentions []struct {
		Name        string `json:"name"`
		Indices     []int `json:"indices"`
		Screen_name string `json:"screen_name"`
		Id          int64 `json:"id"`
		Id_str      string `json:"id_str"`
	} `json:"user_mentions"`
}
