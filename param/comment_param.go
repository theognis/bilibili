package param

import "bilibili/model"

type Comment struct {
	Id      int64
	User    model.SpaceUserinfo
	VideoId int64
	Value   string
	Time    string
	Likes   int64
}
