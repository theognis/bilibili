package model

import "time"

type Userinfo struct {
	Uid             int64
	Username        string
	Password        string
	Email           string
	Phone           string
	Salt            string
	RegDate         time.Time
	LastCheckInDate string
	Statement       string
	Coins           int64
	Exp             int64
}
