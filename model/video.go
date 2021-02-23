package model

import "time"

type Video struct {
	Id          int64
	Title       string
	Channel     string
	Description string
	Video       string
	Cover       string
	Author      int64
	Time        time.Time
	Views       int64
	Likes       int64
	Coins       int64
	Saves       int64
	Shares      int64
	Length      string
}

type VideoWithUserModel struct {
	User        SpaceUserinfo
	Title       string
	Channel     string
	Description string
	Video       string
	Cover       string
	Author      int64
	Time        time.Time
	Views       int64
	Likes       int64
	Coins       int64
	Saves       int64
	Shares      int64
	Length      string
}
