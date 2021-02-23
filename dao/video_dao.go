package dao

import (
	"bilibili/model"
	"bilibili/param"
	"database/sql"
	"time"
)

type VideoDao struct {
	*sql.DB
}

//通过up主uid找up投稿视频的model
func (dao *VideoDao) QueryPostedVideoModelByAuthorUid(uid int64) ([]model.Video, error) {
	var videoSlice []model.Video

	stmt, err := dao.DB.Prepare(`SELECT av, title, channel, description, video_url, cover_url, author_uid, time, views, likes, coins, saves, shares, length FROM video_info WHERE author_uid = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var videoModel model.Video
		err = rows.Scan(&videoModel.Id, &videoModel.Title, &videoModel.Channel, &videoModel.Description, &videoModel.Video, &videoModel.Cover, &videoModel.Author, &videoModel.Time, &videoModel.Views, &videoModel.Likes, &videoModel.Coins, &videoModel.Saves, &videoModel.Shares, &videoModel.Length)
		if err != nil {
			return nil, err
		}

		videoSlice = append(videoSlice, videoModel)
	}

	return videoSlice, nil
}

//查询用户投币过的视频av号
func (dao *VideoDao) QueryCoinsByUid(uid int64) ([]int64, error) {
	var avSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_coin WHERE uid = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		avSlice = append(avSlice, av)
	}

	return avSlice, nil
}

//查询用户收藏过的视频av号
func (dao *VideoDao) QuerySaveByUid(uid int64) ([]int64, error) {
	var avSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_save WHERE uid = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		avSlice = append(avSlice, av)
	}

	return avSlice, nil
}

//查询用户点赞过的视频av号
func (dao *VideoDao) QueryLikesByUid(uid int64) ([]int64, error) {
	var avSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_like WHERE uid = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		avSlice = append(avSlice, av)
	}

	return avSlice, nil
}

//查询点赞过视频的uid
func (dao *VideoDao) QueryLikesByAv(av int64) ([]int64, error) {
	var uidSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT uid FROM video_like WHERE av = ?`)
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
		var uid int64
		err = rows.Scan(&uid)
		if err != nil {
			return nil, err
		}

		uidSlice = append(uidSlice, uid)
	}

	return uidSlice, nil
}

func (dao *VideoDao) QueryAvByLabel(label string) ([]int64, error) {
	var avSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_label WHERE video_label = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(label)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		avSlice = append(avSlice, av)
	}

	return avSlice, nil
}

func (dao *VideoDao) QueryLabel(av int64) ([]string, error) {
	var labelSlice []string

	stmt, err := dao.DB.Prepare(`SELECT video_label FROM video_label WHERE av = ?`)
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
		var label string
		err = rows.Scan(&label)
		if err != nil {
			return nil, err
		}

		labelSlice = append(labelSlice, label)
	}

	return labelSlice, nil
}

func (dao *VideoDao) QueryDanmaku(av int64) ([]param.ParamDanmaku, error) {
	var danmakuSlice []param.ParamDanmaku

	stmt, err := dao.DB.Prepare(`SELECT did, av, uid, value, color, type, time, location FROM video_danmaku WHERE av = ?`)
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
		var danmaku param.ParamDanmaku
		var Time time.Time
		err = rows.Scan(&danmaku.Id, &danmaku.VideoId, &danmaku.UserId, &danmaku.Value, &danmaku.Color, &danmaku.Type, &Time, &danmaku.Location)
		if err != nil {
			return nil, err
		}

		danmaku.Time = Time.Format("2006-01-02 15:04:05")

		danmakuSlice = append(danmakuSlice, danmaku)
	}

	return danmakuSlice, nil
}

func (dao *VideoDao) QueryAvSlice() ([]int64, error) {
	var randAvSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_info ORDER BY rand()`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		randAvSlice = append(randAvSlice, av)
	}

	return randAvSlice, nil
}

func (dao *VideoDao) QueryAvSliceByAuthor(uid int64) ([]int64, error) {
	var randAvSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_info WHERE author_uid = ? ORDER BY rand()`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		randAvSlice = append(randAvSlice, av)
	}

	return randAvSlice, nil
}

func (dao *VideoDao) QuerySameChannelAvSlice(channel string) ([]int64, error) {
	var randAvSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_info WHERE channel LIKE ? ORDER BY rand()`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(channel)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		randAvSlice = append(randAvSlice, av)
	}

	return randAvSlice, nil
}

//按收视率排行的特定分区视频10个，返回AV号
func (dao *VideoDao) QueryRankChannel(channel string) ([]int64, error) {
	var randAvSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_info WHERE channel LIKE ? ORDER BY views DESC LIMIT 10`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(channel)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		randAvSlice = append(randAvSlice, av)
	}

	return randAvSlice, nil
}

//按收视率排行的”资讯“8个,返回AV号
func (dao *VideoDao) QueryRankInfo() ([]int64, error) {
	var rankAvSlice []int64

	stmt, err := dao.DB.Prepare(`select av from video_info where av in (select av from video_label where video_label = ?) order by views desc limit 10;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query("资讯")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		rankAvSlice = append(rankAvSlice, av)
	}

	return rankAvSlice, nil
}

//随机特定分区视频8个，返回AV号
func (dao *VideoDao) QueryRandomChannel(channel string) ([]int64, error) {
	var randAvSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_info WHERE channel LIKE ? ORDER BY rand() LIMIT 8`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(channel)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		randAvSlice = append(randAvSlice, av)
	}

	return randAvSlice, nil
}

//随机”资讯“8个,返回AV号
func (dao *VideoDao) QueryRandomInfo() ([]int64, error) {
	var randAvSlice []int64

	stmt, err := dao.DB.Prepare(`SELECT av FROM video_label WHERE video_label = ? ORDER BY rand() LIMIT 8`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query("资讯")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		randAvSlice = append(randAvSlice, av)
	}

	return randAvSlice, nil
}

func (dao *VideoDao) QueryByAv(av int64) (model.Video, error) {
	videoModel := model.Video{}

	stmt, err := dao.DB.Prepare(`SELECT av, title, channel, description, video_url, cover_url, author_uid, time, views, likes, coins, saves, shares, length FROM video_info WHERE av = ?`)
	defer stmt.Close()

	if err != nil {
		return videoModel, err
	}

	row := stmt.QueryRow(av)

	err = row.Scan(&videoModel.Id, &videoModel.Title, &videoModel.Channel, &videoModel.Description, &videoModel.Video, &videoModel.Cover, &videoModel.Author, &videoModel.Time, &videoModel.Views, &videoModel.Likes, &videoModel.Coins, &videoModel.Saves, &videoModel.Shares, &videoModel.Length)
	if err != nil {
		return videoModel, err
	}

	return videoModel, nil
}

func (dao *VideoDao) UpdateViews(av int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE video_info SET views = views + 1 WHERE av = ?`)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(av)
	if err != nil {
		return err
	}

	return nil
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

//coin表插入一条记录
func (dao *VideoDao) InsertCoin(av, uid int64) error {
	stmt, err := dao.DB.Prepare(`INSERT INTO video_coin (av, uid) VALUES (?, ?)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(av, uid)
	return err
}

//save表插入一条记录
func (dao *VideoDao) InsertSave(av, uid int64) error {
	stmt, err := dao.DB.Prepare(`INSERT INTO video_save (av, uid) VALUES (?, ?)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(av, uid)
	return err
}

//like表插入一条记录
func (dao *VideoDao) InsertLike(av, uid int64) error {
	stmt, err := dao.DB.Prepare(`INSERT INTO video_like (av, uid) VALUES (?, ?)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(av, uid)
	return err
}

func (dao *VideoDao) InsertDanmaku(danmakuModel model.Danmaku) (int64, error) {
	stmt, err := dao.DB.Prepare(`INSERT INTO video_danmaku (av, uid, value, color, type, time, location) VALUES (?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(danmakuModel.Av, danmakuModel.Uid, danmakuModel.Value, danmakuModel.Color, danmakuModel.Type, danmakuModel.Time, danmakuModel.Location)

	stmt.Close()

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (dao *VideoDao) InsertVideo(video model.Video) (int64, error) {
	stmt, err := dao.DB.Prepare(`INSERT INTO video_info (title, channel, description, video_url, cover_url, author_uid, time, length) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(video.Title, video.Channel, video.Description, video.Video, video.Cover, video.Author, video.Time, video.Length)

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

func (dao *VideoDao) DeleteSave(av, uid int64) error {
	stmt, err := dao.DB.Prepare(`DELETE FROM video_save WHERE (av = ? and uid = ?)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(av, uid)
	return err
}

func (dao *VideoDao) DeleteLike(av, uid int64) error {
	stmt, err := dao.DB.Prepare(`DELETE FROM video_like WHERE (av = ? and uid = ?)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(av, uid)
	return err
}

func (dao *VideoDao) UpdateSave(av, num int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE video_info SET saves = saves + ? WHERE av = ?`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(num, av)
	if err != nil {
		return err
	}

	return nil
}

func (dao *VideoDao) UpdateLike(av, num int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE video_info SET likes = likes + ? WHERE av = ?`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(num, av)
	if err != nil {
		return err
	}

	return nil
}

func (dao *VideoDao) UpdateShare(av int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE video_info SET shares = shares + 1 WHERE av = ?`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(av)
	if err != nil {
		return err
	}

	return nil
}

func (dao *VideoDao) UpdateCoin(av int64) error {
	stmt, err := dao.DB.Prepare(`UPDATE video_info SET coins = coins + 1 WHERE av = ?`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(av)
	if err != nil {
		return err
	}

	return nil
}

func (dao *VideoDao) Search(keywords string) ([]int64, error) {
	var avSlice []int64

	keywords = "%" + keywords + "%"
	stmt, err := dao.DB.Prepare(`SELECT av FROM video_info WHERE title LIKE ? ORDER BY views DESC`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(keywords)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var av int64
		err = rows.Scan(&av)
		if err != nil {
			return nil, err
		}

		avSlice = append(avSlice, av)
	}

	return avSlice, nil
}
