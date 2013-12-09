package anaconda

type TwitterUser struct {
	Statuses_count                     *int64 `json:"statuses_count"`
	Contributors_enabled               *bool `json:"contributors_enabled"`
	Friends_count                      *int64 `json:"friends_count"`
	Geo_enabled                        *bool `json:"geo_enabled"`
	Description                        *string `json:"description"`
	Profile_sidebar_border_color       *string `json:"profile_sidebar_border_color"`
	Screen_name                        *string `json:"screen_name"`
	Listed_count                       *int64 `json:"listed_count"`
	Followers_count                    *int64 `json:"followers_count"`
	Location                           *string `json:"location"`
	Profile_background_image_url       *string `json:"profile_background_image_url"`
	Name                               *string `json:"name"`
	Default_profile_image              *bool `json:"default_profile_image"`
	Profile_image_url_https            *string `json:"profile_image_url_https"`
	Notifications                      *bool `json:"notifications"`
	Protected                          *bool `json:"protected"`
	Id_str                             *string `json:"id_str"`
	Profile_background_color           *string `json:"profile_background_color"`
	Created_at                         *string `json:"created_at"`
	Default_profile                    *bool `json:"default_profile"`
	Url                                *string `json:"url"`
	Id                                 *int64 `json:"id"`
	Verified                           *bool `json:"verified"`
	Profile_link_color                 *string `json:"profile_link_color"`
	Profile_image_url                  *string `json:"profile_image_url"`
	Profile_use_background_image       *bool `json:"profile_use_background_image"`
	Favourites_count                   *int64 `json:"favourites_count"`
	Profile_background_image_url_https *string `json:"profile_background_image_url_https"`
	Profile_sidebar_fill_color         *string `json:"profile_sidebar_fill_color"`
	Is_translator                      *bool `json:"is_translator"`
	Follow_request_sent                *bool `json:"follow_request_sent"`
	Following                          *bool `json:"following"`
	Profile_background_tile            *bool `json:"profile_background_tile"`
	Show_all_inline_media              *bool `json:"show_all_inline_media"`
	Profile_text_color                 *string `json:"profile_text_color"`
	Lang                               *string `json:"lang"`
	Entities                           *ITwitterEntities `json:"entities"`
}
