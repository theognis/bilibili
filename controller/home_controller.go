package controller

import (
	"bilibili/model"
	"bilibili/service"
	"bilibili/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func (h *HomeController) Router(engine *gin.Engine) {
	engine.GET("/api/home/sections", h.getSections)
	engine.GET("/api/home/search", h.search)
}

func (h *HomeController) search(ctx *gin.Context) {
	keywords := ctx.Query("keywords")

	if keywords == "" {
		tool.Failed(ctx, "搜索内容不可为空")
		return
	}

	hs := service.HomeService{}
	videoWithUserSlice, err := hs.Search(keywords)
	if err != nil {
		fmt.Println("SearchErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if videoWithUserSlice == nil {
		videoWithUserSlice = []model.VideoWithUserModel{}
	}

	tool.Success(ctx, videoWithUserSlice)
}

func (h *HomeController) getSections(ctx *gin.Context) {
	typeSlice := []string{"07%", "16%", "13%", "06%", "12%", "02%", "04%", "10%", "01%", "09%", "11%", "14%", "08%", "information", "03%", "18%", "17%", "05%", "15%"}
	hs := service.HomeService{}
	var Date []model.Section

	for i, channelType := range typeSlice {
		fmt.Println(i)
		//特判资讯
		var section model.Section
		if channelType == "information" {
			randSlice, rankSlice, err := hs.GetInformation()
			if err != nil {
				fmt.Println("GetInformationErr: ", err)
				tool.Failed(ctx, "服务器错误")
				return
			}
			section.List = randSlice
			section.Rank = rankSlice

			if section.List == nil {
				section.List = []model.VideoWithUserModel{}
			}

			if section.Rank == nil {
				section.Rank = []model.VideoWithUserModel{}
			}

			Date = append(Date, section)
			continue
		}
		randSlice, rankSlice, err := hs.GetChannelVideo(channelType)
		if err != nil {
			fmt.Println("GetChannelVideoErr: ", err)
			tool.Failed(ctx, "服务器错误")
			return
		}
		section.List = randSlice
		section.Rank = rankSlice

		if section.List == nil {
			section.List = []model.VideoWithUserModel{}
		}

		if section.Rank == nil {
			section.Rank = []model.VideoWithUserModel{}
		}

		Date = append(Date, section)

	}

	fmt.Println(Date)
	tool.Success(ctx, Date)
}
