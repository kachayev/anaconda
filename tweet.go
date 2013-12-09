package anaconda

type Tweet struct {
	Source        string `json:"source"`
	Id            int64 `json:"id"`
	Retweeted     bool `json:"retweeted"`
	Favorited     bool `json:"favorited"`
	User          TwitterUser `json:"user"`
	Truncated     bool `json:"truncated"`
	Text          string `json:"text"`
	Retweet_count int64 `json:"retweet_count"`
	Id_str        string `json:"id_str"`
	Created_at    string `json:"created_at"`
	Entities      TwitterEntities `json:"entities"`
}
