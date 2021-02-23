package model

type SpaceUserinfo struct {
	Avatar     string
	Uid        int64
	Username   string
	RegDate    string
	Statement  string
	Exp        int64
	Coins      int64
	BCoins     int64
	Birthday   string
	Gender     string
	Videos     []Video
	Saves      []Video
	Followers  int64
	Followings int64
	TotalViews int64
	TotalLikes int64
}
