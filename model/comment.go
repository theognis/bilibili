package model

type Comment struct {
	Id      int64
	VideoId int64
	UserId  int64
	Value   string
	Time    string
	Likes   int64
}
