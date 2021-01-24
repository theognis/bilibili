package controller

import (
	"bilibili/model"
	"bilibili/param"
	"bilibili/service"
	"bilibili/tool"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type UserController struct {

}

func (u *UserController) Router(engine *gin.Engine) {
	engine.POST("/api/user/register", u.register)
	engine.POST("/api/verify/phone", u.sendSms)
	engine.POST("/api/user/username", u.judgeUsername)
	engine.POST("/api/user/password", u.judgePassword)
	engine.POST("/api/user/phone", u.judgePhone)
	engine.POST("/api/user/login", u.login)
}

func (u *UserController) login(ctx *gin.Context) {
	loginName := ctx.PostForm("loginName")
	password := ctx.PostForm("password")
	us := service.UserService{}
	gs := service.GeneralService{}


	userinfo, ok, err := us.Login(loginName, password)
	if err != nil {
		fmt.Println("loginErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if ok == false {
		tool.Failed(ctx, "密码不匹配")
		return
	}

	//创建token， 有效期两分钟
	tokenString, err := gs.CreateToken(userinfo, 120)
	if err != nil {
		fmt.Println("CreateTokenErr:", err)
		tool.Failed(ctx, "系统错误")
		return
	}

	//创建一个refresh token有效期一周
	refreshToken, err := gs.CreateToken(userinfo, 604800)
	if err != nil {
		fmt.Println("CreateRefreshTokenErr:", err)
		tool.Failed(ctx, "系统错误")
		return
	}

	ctx.JSON(200, gin.H{
		"status": "0",
		"data": "登录成功",
		"token": tokenString,
		"refreshToken": refreshToken,
	})
}

func (u *UserController) judgePassword(ctx *gin.Context) {
	password := ctx.PostForm("password")

	if len(password) < 6 {
		tool.Failed(ctx, " 密码不能小于6个字符")
		return
	}

	if len(password) > 16 {
		tool.Failed(ctx, " 密码不能大于16个字符")
		return
	}

	tool.Success(ctx, "密码可以使用")
}

func (u *UserController) judgePhone(ctx *gin.Context) {
	phone := ctx.PostForm("phone")

	if len(phone) != 11 {
		tool.Failed(ctx, "手机号格式不正确")
		return
	}

	us := service.UserService{}
	flag := us.JudgePhone(phone)
	if flag == false {
		tool.Success(ctx, "该电话可以使用")
		return
	}

	tool.Failed(ctx, "该电话已经被注册")
}

//检验用户名是否重复，是否符合规范
func (u *UserController) judgeUsername(ctx *gin.Context) {
	username := ctx.PostForm("username")

	if username == "" || len(username) > 14 {
		tool.Failed(ctx, "用户名不符合规范")
		return
	}

	us := service.UserService{}
	flag := us.JudgeUsername(username)
	if flag == false {
		tool.Success(ctx, "该用户名可以使用")
		return
	}

	tool.Failed(ctx, "用户名重复")
}

//发送短信验证码
func (u *UserController) sendSms(ctx *gin.Context) {
	phone := ctx.PostForm("phone")

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
		tool.Failed(ctx, "服务器错误")
		fmt.Println("SetRedisErr: ", err)
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
	//givenCode := userParam.VerifyCode

	us := service.UserService{}
	//flag, err := us.JudgePhoneCode(ctx, phone, givenCode)
	//if err != nil {
	//	tool.Failed(ctx, "服务器错误")
	//	fmt.Println("JudgePhoneCodeErr: ", err)
	//	return
	//}
	//
	//if flag == false {
	//	tool.Failed(ctx, "验证码错误")
	//	return
	//}

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
	err = us.RegisterModelIn(user)
	if err != nil {
		fmt.Println("RegisterModelInErr: ", err)
		tool.Failed(ctx, "服务器错误")
	}

	tool.Success(ctx, "注册成功！")
}