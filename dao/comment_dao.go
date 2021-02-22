package dao

import (
	"bilibili/model"
	"database/sql"
	"time"
)

type CommentDao struct {
	*sql.DB
}

func (dao *CommentDao) QueryByAv(av int64) ([]model.Comment, error) {
	var commentSlice []model.Comment
	var commentTime time.Time

	stmt, err := dao.DB.Prepare(`SELECT id, av, uid, value, comment_time, likes FROM comment WHERE av = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(av)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var commentModel model.Comment
		err = rows.Scan(&commentModel.Id, &commentModel.VideoId, &commentModel.UserId, &commentModel.Value, &commentTime, &commentModel.Likes)
		if err != nil {
			return nil, err
		}

		commentModel.Time = commentTime.Format("2006-01-02 15:04:05")

		commentSlice = append(commentSlice, commentModel)
	}

	return commentSlice, nil
}

func (dao *CommentDao) InsertComment(comment model.Comment) error {
	timeNow := time.Now()
	stmt, err := dao.DB.Prepare(`INSERT INTO comment (av, uid, value, comment_time) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(comment.VideoId, comment.UserId, comment.Value, timeNow)

	stmt.Close()

	return err
}
