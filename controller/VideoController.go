package controller

import (
	"bilibili/model"
	"bilibili/param"
	"bilibili/service"
	"bilibili/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type VideoController struct {
}

func (v *VideoController) Router(engine *gin.Engine) {
	engine.POST("/api/video/video", v.postVideo)
	engine.POST("/api/video/danmaku", v.postDanmaku)
}

func (v *VideoController) postDanmaku(ctx *gin.Context) {
	var danmakuParam param.PostDanmakuParam
	err := ctx.BindJSON(&danmakuParam)
	if err != nil {
		fmt.Println("BindJsonErr: ", err)
		tool.Failed(ctx, "参数无效")
		return
	}

	if danmakuParam.Token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	gs := service.TokenService{}
	//解析token
	clams, err := gs.ParseToken(danmakuParam.Token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo

	if danmakuParam.Value == "" {
		tool.Failed(ctx, "弹幕不可为空")
		return
	}

	if len(danmakuParam.Value) > 100 {
		tool.Failed(ctx, "弹幕过长")
		return
	}

	if len(danmakuParam.Color) != 6 {
		fmt.Println("parseColorErr: ", err)
		tool.Failed(ctx, "参数无效")
		return
	}

	_, err = strconv.ParseInt(danmakuParam.Color, 16, 64)
	if err != nil {
		fmt.Println("parseColorErr: ", err)
		tool.Failed(ctx, "参数无效")
		return
	}

	//type判断
	if danmakuParam.Type != "scroll" && danmakuParam.Type != "top" && danmakuParam.Type != "bottom" {
		fmt.Println("TypeErr")
		tool.Failed(ctx, "参数无效")
		return
	}

	//location判断
	if danmakuParam.Location < 0 || danmakuParam.Location >= 9999 {
		fmt.Println("LocationErr")
		tool.Failed(ctx, "参数无效")
		return
	}

	//av号判断
	vs := service.VideoService{}
	flag, err = vs.JudgeAv(danmakuParam.Av)
	if err != nil {
		fmt.Println("JudgeAvErr:", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if flag == false {
		tool.Failed(ctx, "参数无效")
		fmt.Println("AvNumNil")
		return
	}

	var danmakuModel model.Danmaku
	danmakuModel.Av = danmakuParam.Av
	danmakuModel.Location = danmakuParam.Location
	danmakuModel.Type = danmakuParam.Type
	danmakuModel.Color = danmakuParam.Color
	danmakuModel.Value = danmakuParam.Value
	danmakuModel.Uid = userinfo.Uid
	danmakuModel.Time = time.Now()

	err = vs.InsertDanmaku(danmakuModel)
	if err != nil {
		fmt.Println("InsertDanmakuErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	tool.Success(ctx, "")
}

func (v *VideoController) postVideo(ctx *gin.Context) {
	token := ctx.PostForm("token")

	if token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	gs := service.TokenService{}
	//解析token
	clams, err := gs.ParseToken(token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	//	userinfo := clams.Userinfo

	//视频文件判断相关
	videoFile, videoHeader, err := ctx.Request.FormFile("video")
	if err != nil {
		tool.Failed(ctx, "视频上传失败")
		return
	}

	//视频大小判断
	if videoHeader.Size > (2048 << 20) {
		tool.Failed(ctx, "视频体积不可大于 2GB")
		return
	}

	if videoHeader.Size == 0 {
		tool.Failed(ctx, "视频不可为空")
		return
	}

	//视频格式判断
	videoExtension := tool.GetExtension(videoHeader.Filename)
	videoExtension = strings.ToLower(videoExtension)
	if videoExtension != "flv" && videoExtension != "mp4" {
		tool.Failed(ctx, "视频格式无效")
		return
	}

	//封面文件判断相关
	coverFile, coverHeader, err := ctx.Request.FormFile("cover")
	if err != nil {
		tool.Failed(ctx, "封面上传失败")
		return
	}

	//封面大小判断
	if coverHeader.Size > (2 << 20) {
		tool.Failed(ctx, "封面体积不可大于 2MB")
		return
	}

	if coverHeader.Size == 0 {
		tool.Failed(ctx, "封面不可为空")
		return
	}

	//视频格式判断
	coverExtension := tool.GetExtension(coverHeader.Filename)
	coverExtension = strings.ToLower(coverExtension)
	if coverExtension != "png" && videoExtension != "jpg" {
		tool.Failed(ctx, "封面格式无效")
		return
	}

	//标题相关
	title := ctx.PostForm("title")
	if title == "" {
		tool.Failed(ctx, "标题不可为空")
		return
	}

	if len(title) > 80 {
		tool.Failed(ctx, "标题过长")
		return
	}

	//频道相关
	channel := ctx.PostForm("channel")

	channelSlice := []string{"0101", "0102", "0103", "0104", "0105", "0106", "0107", "0201", "0202", "0203", "0204", "0205", "0206", "0207", "0208", "0301", "0302", "0401", "0402", "0403", "0404", "0405", "0406", "0501", "0502", "0503", "0504", "0601", "0602", "0603", "0604", "0605", "0606", "0607", "0608", "0609", "0701", "0702", "0703", "0704", "0705", "0706", "0801", "0802", "0803", "0804", "0805", "0901", "0902", "0903", "0904", "0905", "1001", "1002", "1003", "1004", "1101", "1102", "1103", "1104", "1105", "1106", "1201", "1202", "1203", "1204", "1205", "1206", "1301", "1302", "1303", "1304", "1305", "1401", "1402", "1403", "1404", "1405", "1406", "1501", "1502", "1503", "1504", "1601", "1602", "1701", "1702", "1801", "1802", "1803", "1804"}
	flag = false

	for _, channelType := range channelSlice {
		if channelType == channel {
			flag = true
			break
		}
	}

	if flag == false {
		tool.Failed(ctx, "分区无效")
		return
	}

	//简介相关
	description := ctx.PostForm("description")
	if description == "" {
		description = "19260817"
	}

	if len(description) > 250 {
		tool.Failed(ctx, "简介过长")
		return
	}

	//标签相关
	labelStr := ctx.PostForm("label")
	var label []string
	err = json.Unmarshal([]byte(labelStr), &label)
	if err != nil {
		tool.Failed(ctx, "标签无效")
		return
	}

	//切片去重并判断
	var result []string
	tempMap := map[string]byte{}
	for _, e := range label {
		l := len(tempMap)
		tempMap[e] = 0

		if len(tempMap) != l {
			result = append(result, e)
		}
	}

	labelNum := 0
	for _, singleLabel := range result {
		labelNum++
		if len(singleLabel) > 10 {
			tool.Failed(ctx, "标签过长")
			return
		}
	}

	if labelNum == 0 || labelNum > 10 {
		tool.Failed(ctx, "标签无效")
		return
	}

	vs := service.VideoService{}

	//视频入数据库
	var videoInfoModel model.Video
	videoInfoModel.Time = time.Now()
	videoInfoModel.AuthorUid = clams.Userinfo.Uid
	videoInfoModel.Channel = channel
	cfg := tool.GetCfg().Oss
	videoInfoModel.CoverUrl = "nil"
	videoInfoModel.VideoUrl = "nil"
	videoInfoModel.Description = description
	videoInfoModel.Title = title

	av, err := vs.InsertVideo(videoInfoModel)
	if err != nil {
		fmt.Println("InsertVideoErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	videoUrl := cfg.VideosUrl + strconv.FormatInt(av, 10) + "." + videoExtension
	coverUrl := cfg.VideosUrl + strconv.FormatInt(av, 10) + "." + coverExtension

	err = vs.SetUrl(av, videoUrl, coverUrl)
	if err != nil {
		fmt.Println("SetUrlErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	//上传视频
	Os := service.OssService{}

	err = Os.UploadVideoBucket(videoFile, strconv.FormatInt(av, 10)+"."+videoExtension)
	if err != nil {
		fmt.Println("UploadVideoBucketVideoErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	err = Os.UploadVideoBucket(coverFile, strconv.FormatInt(av, 10)+"."+coverExtension)
	if err != nil {
		fmt.Println("UploadVideoBucketCoverErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	//标签入库
	err = vs.InsertLabel(result, av)
	if err != nil {
		fmt.Println("InsertLabelErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	tool.Success(ctx, "")
}
