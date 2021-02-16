package dao

import (
	"bilibili/model"
	"database/sql"
)

type VideoDao struct {
	*sql.DB
}

func (dao *VideoDao) QueryByAv(av int64) (model.Video, error) {
	videoModel := model.Video{}

	stmt, err := dao.DB.Prepare(`SELECT av, title, channel, description, video_url, cover_url, author_uid, time, views, likes, coins, saves, shares FROM video_info WHERE av = ?`)
	defer stmt.Close()

	if err != nil {
		return videoModel, err
	}

	row := stmt.QueryRow(av)

	err = row.Scan(&videoModel.Av, &videoModel.Title, &videoModel.Channel, &videoModel.Description, &videoModel.VideoUrl, &videoModel.CoverUrl, &videoModel.AuthorUid, &videoModel.Time, &videoModel.Views, &videoModel.Likes, &videoModel.Coins, &videoModel.Saves, &videoModel.Shares)
	if err != nil {
		return videoModel, err
	}

	return videoModel, nil
}

func (dao *VideoDao) UpdateUrl(av int64, videoUrl string, coverUrl string) error {
	stmt, err := dao.DB.Prepare(`UPDATE video_info SET video_url = ?, cover_Url = ? WHERE av = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(videoUrl, coverUrl, av)
	if err != nil {
		return err
	}

	return nil
}

func (dao *VideoDao) InsertDanmaku(danmakuModel model.Danmaku) error {
	stmt, err := dao.DB.Prepare(`INSERT INTO video_danmaku (av, uid, value, color, type, time, location) VALUES (?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(danmakuModel.Av, danmakuModel.Uid, danmakuModel.Value, danmakuModel.Color, danmakuModel.Type, danmakuModel.Time, danmakuModel.Location)

	stmt.Close()

	return err
}

func (dao *VideoDao) InsertVideo(video model.Video) (int64, error) {
	stmt, err := dao.DB.Prepare(`INSERT INTO video_info (title, channel, description, video_url, cover_url, author_uid, time) VALUES (?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(video.Title, video.Channel, video.Description, video.VideoUrl, video.CoverUrl, video.AuthorUid, video.Time)

	stmt.Close()
	av, _ := result.LastInsertId()

	return av, err
}

func (dao *VideoDao) InsertLabel(label string, av int64) error {
	stmt, err := dao.DB.Prepare(`INSERT INTO video_label (av, video_label) VALUES (?, ?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(av, label)

	stmt.Close()

	return err
}
