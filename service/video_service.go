package service

import (
	"bilibili/dao"
	"bilibili/model"
	"bilibili/param"
	"bilibili/tool"
	"time"
)

type VideoService struct {
}

func (v *VideoService) AddShare(av int64) error {
	vd := dao.VideoDao{tool.GetDb()}

	err := vd.UpdateShare(av)
	return err
}

func (v *VideoService) GetAvSlice() ([]int64, error) {
	vd := dao.VideoDao{tool.GetDb()}

	avSlice, err := vd.QueryAvSlice()
	if err != nil {
		return nil, err
	}

	return avSlice, nil
}

func (v *VideoService) GetSameUpAvSlice(av int64) ([]int64, error) {
	vd := dao.VideoDao{tool.GetDb()}

	userinfo, err := vd.QueryByAv(av)
	if err != nil {
		return nil, err
	}
	uid := userinfo.Author

	avSlice, err := vd.QueryAvSliceByAuthor(uid)
	if err != nil {
		return nil, err
	}

	return avSlice, nil
}

func (v *VideoService) GetSameChannelAvSlice(av int64) ([]int64, error) {
	vd := dao.VideoDao{tool.GetDb()}

	userinfo, err := vd.QueryByAv(av)
	if err != nil {
		return nil, err
	}

	channel := userinfo.Channel[:2] + "%"
	avSlice, err := vd.QuerySameChannelAvSlice(channel)
	if err != nil {
		return nil, err
	}

	return avSlice, nil
}

//获取一个视频是否被收藏，已被收藏返回true
func (v *VideoService) JudgeSave(uid int64, av int64) (bool, error) {
	vd := dao.VideoDao{tool.GetDb()}

	avSlice, err := vd.QuerySaveByUid(uid)
	if err != nil {
		return false, err
	}

	for _, savedAv := range avSlice {
		if av == savedAv {
			return true, nil
		}
	}

	return false, nil
}

//收藏/取消收藏
func (v *VideoService) PostSave(uid int64, av int64, flag bool) error {
	vd := dao.VideoDao{tool.GetDb()}
	var err error

	if flag == false {
		//此前未收藏
		err = vd.InsertSave(av, uid)
		if err != nil {
			return err
		}

		err = vd.UpdateSave(av, 1)
		if err != nil {
			return err
		}
	} else {
		//此前已收藏
		err = vd.DeleteSave(av, uid)
		if err != nil {
			return err
		}

		err = vd.UpdateSave(av, -1)
		if err != nil {
			return err
		}
	}
	return nil

}
func (v *VideoService) PostCoin(av, uid int64) (bool, error) {
	vd := dao.VideoDao{tool.GetDb()}
	ud := dao.UserDao{tool.GetDb()}

	//获取up uid
	videoInfo, err := vd.QueryByAv(av)
	if err != nil {
		return false, err
	}
	upUid := videoInfo.Author

	if upUid == uid {
		return false, nil
	}
	//up加经验
	err = ud.UpdateExp(upUid, 1)
	if err != nil {
		return false, err
	}

	//获取当日投币所得经验状态
	userinfo, err := ud.QueryByUid(uid)
	if err != nil {
		return false, err
	}

	lastCoinDate := userinfo.LastCoinDate[:10]
	timeNow := time.Now().Format("2006-01-02")

	if (lastCoinDate != timeNow) || ((lastCoinDate == timeNow) && (userinfo.DailyCoin < 5)) {
		//处理经验
		ud.UpdateExp(uid, 10)

		//处理记录
		err = ud.UpdateLastCoinDate(uid)
		if err != nil {
			return false, err
		}

		if lastCoinDate != timeNow {
			//上一次投币时间不是今天， 把当日投币次数改为1
			err = ud.UpdateDailyCoin(uid, 1)
			if err != nil {
				return false, err
			}
		} else {
			//上一次投币时间是今天，把当日投币次数加一
			err = ud.UpdateDailyCoin(uid, -1)
			if err != nil {
				return false, err
			}
		}

	}

	//硬币转账
	err = ud.UpdateCoins(uid, -1)
	if err != nil {
		return false, err
	}

	err = ud.UpdateCoins(upUid, 1)
	if err != nil {
		return false, err
	}

	//添加记录
	err = vd.InsertCoin(av, uid)
	if err != nil {
		return false, err
	}

	//视频信息更新
	err = vd.UpdateCoin(av)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (v *VideoService) SolveLike(flag bool, uid int64, av int64) error {
	vd := dao.VideoDao{tool.GetDb()}
	var err error

	if flag == false {
		//此前未点赞
		err = vd.InsertLike(av, uid)
		if err != nil {
			return err
		}

		err = vd.UpdateLike(av, 1)
		if err != nil {
			return err
		}
	} else {
		//此前已点赞
		err = vd.DeleteLike(av, uid)
		if err != nil {
			return err
		}

		err = vd.UpdateLike(av, -1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *VideoService) GetCoin(av, uid int64) (bool, error) {
	vd := dao.VideoDao{tool.GetDb()}

	avSlice, err := vd.QueryCoinsByUid(uid)
	if err != nil {
		return false, err
	}

	for _, coinedAv := range avSlice {
		if av == coinedAv {
			return true, nil
		}
	}

	return false, nil
}

//获取用户点赞状态，在err为nil的情况下，已经点赞返回true，反正返回false
//考虑到单个视频可能存在大量赞，这里在dao层查询用户点赞的视频，而不是查询点赞过视频的用户，优化性能
func (v *VideoService) GetLike(av, uid int64) (bool, error) {
	vd := dao.VideoDao{tool.GetDb()}

	avSlice, err := vd.QueryLikesByUid(uid)
	if err != nil {
		return false, err
	}

	for _, likedAv := range avSlice {
		if av == likedAv {
			return true, nil
		}
	}

	return false, nil
}

func (v *VideoService) GetVideo(av int64) (model.Video, error) {
	vd := dao.VideoDao{tool.GetDb()}

	videoInfo, err := vd.QueryByAv(av)
	return videoInfo, err
}

func (v *VideoService) GetAvByLabel(label string) ([]int64, error) {
	vd := dao.VideoDao{tool.GetDb()}

	avSlice, err := vd.QueryAvByLabel(label)
	return avSlice, err
}

func (v *VideoService) GetLabel(av int64) ([]string, error) {
	vd := dao.VideoDao{tool.GetDb()}

	labelSlice, err := vd.QueryLabel(av)
	return labelSlice, err
}

func (v *VideoService) GetDanmaku(av int64) ([]param.ParamDanmaku, error) {
	vd := dao.VideoDao{tool.GetDb()}

	danmakuSlice, err := vd.QueryDanmaku(av)
	return danmakuSlice, err

}

func (v *VideoService) AddView(av int64) error {
	vd := dao.VideoDao{tool.GetDb()}
	err := vd.UpdateViews(av)
	return err
}

func (v *VideoService) InsertDanmaku(danmakuModel model.Danmaku) (int64, error) {
	vd := dao.VideoDao{tool.GetDb()}
	did, err := vd.InsertDanmaku(danmakuModel)
	return did, err
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
