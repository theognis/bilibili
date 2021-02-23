package service

import (
	"bilibili/dao"
	"bilibili/model"
	"bilibili/tool"
)

type HomeService struct {
}

func (hs *HomeService) Search(keywords string) ([]model.VideoWithUserModel, error) {
	vd := dao.VideoDao{tool.GetDb()}
	ud := dao.UserDao{tool.GetDb()}
	var result []model.VideoWithUserModel
	var videoWithUser model.VideoWithUserModel

	avSlice, err := vd.Search(keywords)
	if err != nil {
		return nil, err
	}

	for _, av := range avSlice {
		video, err := vd.QueryByAv(av)
		if err != nil {
			return nil, err
		}

		userModel, err := ud.QuerySpaceUserinfoByUid(video.Author)
		if err != nil {
			return nil, err
		}

		videoWithUser.User = userModel
		videoWithUser.Id = video.Id
		videoWithUser.Likes = video.Likes
		videoWithUser.Time = video.Time
		videoWithUser.Author = video.Author
		videoWithUser.Channel = video.Channel
		videoWithUser.Video = video.Video
		videoWithUser.Saves = video.Saves
		videoWithUser.Length = video.Length
		videoWithUser.Coins = video.Coins
		videoWithUser.Shares = video.Shares
		videoWithUser.Views = video.Views
		videoWithUser.Description = video.Description
		videoWithUser.Title = video.Title
		videoWithUser.Cover = video.Cover

		result = append(result, videoWithUser)
	}

	return result, nil
}

//获取显示在首页的各个分区视频信息
func (hs *HomeService) GetChannelVideo(channel string) ([]model.Video, []model.Video, error) {
	vd := dao.VideoDao{tool.GetDb()}
	var randSlice, rankSlice []model.Video

	randAvSlice, err := vd.QueryRandomChannel(channel)
	if err != nil {
		return randSlice, rankSlice, err
	}

	for _, randAv := range randAvSlice {
		var videoModel model.Video
		videoModel, err = vd.QueryByAv(randAv)
		if err != nil {
			return randSlice, rankSlice, err
		}

		randSlice = append(randSlice, videoModel)
	}

	rankAvSlice, err := vd.QueryRankChannel(channel)
	if err != nil {
		return randSlice, rankSlice, err
	}

	for _, rankAv := range rankAvSlice {
		var videoModel model.Video
		videoModel, err = vd.QueryByAv(rankAv)
		if err != nil {
			return randSlice, rankSlice, err
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
		return randSlice, rankSlice, err
	}

	for _, randAv := range randAvSlice {
		var videoModel model.Video
		videoModel, err = vd.QueryByAv(randAv)
		if err != nil {
			return randSlice, rankSlice, err
		}

		randSlice = append(randSlice, videoModel)
	}

	rankAvSlice, err := vd.QueryRankInfo()
	if err != nil {
		return randSlice, rankSlice, err
	}

	for _, rankAv := range rankAvSlice {
		var videoModel model.Video
		videoModel, err = vd.QueryByAv(rankAv)
		if err != nil {
			return randSlice, rankSlice, err
		}

		rankSlice = append(rankSlice, videoModel)
	}

	return randSlice, rankSlice, nil
}
