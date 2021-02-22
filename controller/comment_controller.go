package controller

import (
	"bilibili/model"
	"bilibili/service"
	"bilibili/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"unicode/utf8"
)

type CommentController struct {
}

func (c *CommentController) Router(engine *gin.Engine) {
	engine.GET("/api/video/comments", c.getComments)
	engine.POST("/api/video/comment", c.postComment)
}

func (c *CommentController) getComments(ctx *gin.Context) {
	avStr := ctx.Query("video_id")
	if avStr == "" {
		tool.Failed(ctx, "视频ID不可为空")
		return
	}
	avInt, err := strconv.ParseInt(avStr, 10, 64)
	if err != nil {
		fmt.Println("ParseAvStrErr: ", err)
		tool.Failed(ctx, "视频ID无效")
		return
	}

	vs := service.VideoService{}
	flag, err := vs.JudgeAv(avInt)
	if err != nil {
		fmt.Println("JudgeAvErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if flag == false {
		tool.Failed(ctx, "视频ID无效")
		return
	}

	cs := service.CommentService{}
	commentSlice, err := cs.GetCommentSlice(avInt)
	if err != nil {
		fmt.Println("GetCommentSliceErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if commentSlice == nil {
		commentSlice = []model.Comment{}
	}

	tool.Success(ctx, commentSlice)
}

func (c *CommentController) postComment(ctx *gin.Context) {
	avStr := ctx.PostForm("video_id")
	if avStr == "" {
		tool.Failed(ctx, "视频ID不可为空")
		return
	}
	avInt, err := strconv.ParseInt(avStr, 10, 64)
	if err != nil {
		fmt.Println("ParseAvStrErr: ", err)
		tool.Failed(ctx, "视频ID无效")
		return
	}

	vs := service.VideoService{}
	flag, err := vs.JudgeAv(avInt)
	if err != nil {
		fmt.Println("JudgeAvErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if flag == false {
		tool.Failed(ctx, "视频ID无效")
		return
	}

	token := ctx.PostForm("token")

	if token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	gs := service.TokenService{}
	//解析token
	clams, err := gs.ParseToken(token)
	flag = tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo

	comment := ctx.PostForm("comment")
	if comment == "" {
		tool.Failed(ctx, "评论内容不可为空")
		return
	}

	if utf8.RuneCountInString(comment) > 1024 {
		tool.Failed(ctx, "评论内容过长")
		return
	}

	var commentModel model.Comment
	commentModel.Value = comment
	commentModel.UserId = userinfo.Uid
	commentModel.VideoId = avInt

	cs := service.CommentService{}

	id, err := cs.PostComment(commentModel)
	if err != nil {
		fmt.Println("postCommentErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	commentModel.Time = time.Now().Format("2006-01-02 15:04:05")
	commentModel.Likes = 0
	commentModel.Id = id

	tool.Success(ctx, commentModel)
}
