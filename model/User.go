package model

import "time"

type Userinfo struct {
	Uid             int64
	Username        string
	Password        string
	Email           string
	Phone           string
	Salt            string
	Avatar          string
	RegDate         time.Time
	Birthday        time.Time
	LastCheckInDate string
	Statement       string
	Coins           int64
	Exp             int64
	BCoins          int64
	Gender          string
}
