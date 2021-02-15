package service

import (
	"bilibili/dao"
	"bilibili/model"
	"bilibili/tool"
)

type VideoService struct {
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
