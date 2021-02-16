package model

import "time"

type Video struct {
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
	Shares      int64
}
