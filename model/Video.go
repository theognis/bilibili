package model

import "time"

type Userinfo struct {
	Av          int64
	Title       string
	Channel     string
	Description string
	VideoUrl    string
	CoverUrl    string
	AuthorUid   int64
	Time        time.Time
	Views       int64
	Likes       int64
	Coins       int64
	Saves       int64
	shares      int64
}
