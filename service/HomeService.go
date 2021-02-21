package service

import (
	"bilibili/dao"
	"bilibili/model"
	"bilibili/tool"
)

type HomeService struct {
}

//获取显示在首页的各个分区视频信息
func (hs *HomeService) GetChannelVideo(channel string) ([]model.Video, []model.Video, error) {
	vd := dao.VideoDao{tool.GetDb()}
	var randSlice, rankSlice []model.Video

	randAvSlice, err := vd.QueryRandomChannel(channel)
	if err != nil {
		return nil, nil, err
	}

	for _, randAv := range randAvSlice {
		var videoModel model.Video
		videoModel, err = vd.QueryByAv(randAv)
		if err != nil {
			return nil, nil, err
		}

		randSlice = append(randSlice, videoModel)
	}

	rankAvSlice, err := vd.QueryRankChannel(channel)
	if err != nil {
		return nil, nil, err
	}

	for _, rankAv := range rankAvSlice {
		var videoModel model.Video
		videoModel, err = vd.QueryByAv(rankAv)
		if err != nil {
			return nil, nil, err
		}

		rankSlice = append(rankSlice, videoModel)
	}

	return randSlice, rankSlice, nil
}

//获取显示在首页的”资讯“视频信息
func (hs *HomeService) GetInformation() ([]model.Video, []model.Video, error) {
	vd := dao.VideoDao{tool.GetDb()}
	var randSlice, rankSlice []model.Video

	randAvSlice, err := vd.QueryRandomInfo()
	if err != nil {
		return nil, nil, err
	}

	for _, randAv := range randAvSlice {
		var videoModel model.Video
		videoModel, err = vd.QueryByAv(randAv)
		if err != nil {
			return nil, nil, err
		}

		randSlice = append(randSlice, videoModel)
	}

	rankAvSlice, err := vd.QueryRankInfo()
	if err != nil {
		return nil, nil, err
	}

	for _, rankAv := range rankAvSlice {
		var videoModel model.Video
		videoModel, err = vd.QueryByAv(rankAv)
		if err != nil {
			return nil, nil, err
		}

		rankSlice = append(rankSlice, videoModel)
	}

	return randSlice, rankSlice, nil
}
