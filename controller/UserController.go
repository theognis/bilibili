package controller

import (
	"bilibili/service"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"bilibili/param"
	"bilibili/tool"
	"bilibili/model"
	"strconv"
	"time"
)

type UserController struct {

}

func (u *UserController) Router(engine *gin.Engine) {
	engine.POST("/api/user/register", u.register)
	engine.POST("/api/verify/phone", u.sendSms)
}

//发送短信验证码
func (u *UserController) sendSms(ctx *gin.Context) {
	phone, exist := ctx.GetQuery("phone")
	if !exist {
		tool.Failed(ctx, "参数解析失败")
		return
	}

	us := service.UserService{}
	verifyCode, err := us.SendCodeByPhone(phone)
	if err != nil {
		tool.Failed(ctx, "SendCodeByPhone: "+ err.Error())
		fmt.Println("SendCodeByPhoneErr")
		return
	}
	if verifyCode == "" {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("SendCodeByPhoneErr")
		return
	}

	//把验证码放到redis中
	err = us.PhoneCodeIn(ctx, phone, verifyCode)
	if err != nil {
		tool.Failed(ctx, err)
		fmt.Println("SetRedisErr")
		return
	}

	tool.Success(ctx, "发送验证码成功")
}

func (u *UserController) register(ctx *gin.Context) {
	//获取并解析用户表单
	var userParam param.UserParam
	err := ctx.ShouldBind(&userParam)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
		return
	}

	//判断验证码是否正确
	phone := userParam.Phone
	givenCode := userParam.VerifyCode

	us := service.UserService{}
	flag, err := us.JudgePhoneCode(ctx, phone, givenCode)
	if err != nil {
		tool.Failed(ctx, err)
		fmt.Println("JudgePhoneCodeErr")
		return
	}

	if flag == false {
		tool.Failed(ctx, "验证码错误")
		return
	}

	//创建实体
	var user model.Userinfo
	user.RegDate = time.Now()
	user.Phone = phone
	user.Username = userParam.Username
	//撒盐对md5加密，数据库中非明文存储
	user.Salt = strconv.FormatInt(time.Now().Unix(), 10)
	m5 := md5.New()
	m5.Write([]byte(userParam.Pwd))
	m5.Write([]byte(user.Salt))
	st := m5.Sum(nil)
	user.Password = hex.EncodeToString(st)

	//放入mysql
}