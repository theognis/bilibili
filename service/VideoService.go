package service

import (
	"bilibili/dao"
	"bilibili/model"
	"bilibili/param"
	"bilibili/tool"
)

type VideoService struct {
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
