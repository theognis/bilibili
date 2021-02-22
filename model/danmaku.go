package model

import "time"

type Danmaku struct {
	Did      int64
	Av       int64
	Uid      int64
	Value    string
	Color    string
	Type     string
	Time     time.Time
	Location int64
}
