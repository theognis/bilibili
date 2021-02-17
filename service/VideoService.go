package service

import (
	"bilibili/dao"
	"bilibili/model"
	"bilibili/tool"
)

type VideoService struct {
}

func (v *VideoService) GetDanmaku(av int64) ([]model.Danmaku, error) {
	vd := dao.VideoDao{tool.GetDb()}

	danmakuSlice, err := vd.QueryDanmaku(av)
	return danmakuSlice, err

}

func (v *VideoService) InsertDanmaku(danmakuModel model.Danmaku) error {
	vd := dao.VideoDao{tool.GetDb()}
	err := vd.InsertDanmaku(danmakuModel)
	return err
}

//判断av号对应的视频是否存在，存在则返回true
func (v *VideoService) JudgeAv(av int64) (bool, error) {
	vd := dao.VideoDao{tool.GetDb()}

	_, err := vd.QueryByAv(av)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (v *VideoService) SetUrl(av int64, videoUrl string, coverUrl string) error {
	vd := dao.VideoDao{tool.GetDb()}

	err := vd.UpdateUrl(av, videoUrl, coverUrl)
	return err
}

func (v *VideoService) InsertLabel(labelSlice []string, av int64) error {
	vd := dao.VideoDao{tool.GetDb()}

	for _, label := range labelSlice {
		err := vd.InsertLabel(label, av)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *VideoService) InsertVideo(video model.Video) (int64, error) {
	vd := dao.VideoDao{tool.GetDb()}

	av, err := vd.InsertVideo(video)
	return av, err
}
