package service

import (
	"bilibili/dao"
	"bilibili/model"
	"bilibili/tool"
)

type CommentService struct {
}

func (c *CommentService) GetCommentSlice(av int64) ([]model.Comment, error) {
	cd := dao.CommentDao{tool.GetDb()}

	commentSlice, err := cd.QueryByAv(av)
	return commentSlice, err
}

func (c *CommentService) PostComment(comment model.Comment) error {
	cd := dao.CommentDao{tool.GetDb()}

	err := cd.InsertComment(comment)
	return err
}
