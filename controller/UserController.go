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
	"strings"
	"time"
)

type UserController struct {
}

func (u *UserController) Router(engine *gin.Engine) {
	engine.GET("/api/user/info/self", u.getSelfInfo)
	engine.GET("/api/check/username", u.judgeUsername)
	engine.GET("/api/check/phone", u.judgePhone)
	engine.POST("/api/user/register", u.register)
	engine.POST("/api/verify/sms/register", u.sendSmsRegister)
	engine.POST("/api/verify/sms/general", u.sendSms)
	engine.POST("/api/user/login", u.login)
	engine.POST("/api/verify/email", u.sendEmailCode)
	engine.PUT("/api/user/phone", u.changePhone)
	engine.PUT("/api/user/email", u.changeEmail)
	engine.PUT("/api/user/statement", u.changeStatement)
}

func (u *UserController) sendSmsRegister(ctx *gin.Context) {
	phone := ctx.PostForm("phone")

	if phone == "" {
		tool.Failed(ctx, "手机号不可为空")
	}

	us := service.UserService{}

	flag, err := us.JudgePhone(phone)
	if err != nil {
		fmt.Println("JudgePhoneErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if flag == true {
		tool.Failed(ctx, "手机号已被使用")
		return
	}

	verifyCode, err := us.SendCodeByPhone(phone)
	if err != nil {
		tool.Failed(ctx, "系统错误")
		fmt.Println("SendCodeByPhoneErr")
		return
	}

	if verifyCode == "isv.MOBILE_NUMBER_ILLEGAL" {
		tool.Failed(ctx, "手机号不合法")
		fmt.Println("sendCodeByPhoneErr")
		return
	}

	//把验证码放到redis中
	err = us.VerifyCodeIn(ctx, phone, verifyCode)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("SetRedisErr: ", err)
		return
	}

	tool.Success(ctx, "")
}

func (u *UserController) judgePhone(ctx *gin.Context) {
	phone := ctx.Query("phone")

	if phone == "" {
		tool.Failed(ctx, "请告诉我你的手机号吧")
		return
	}

	us := service.UserService{}
	flag, err := us.JudgePhone(phone)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		return
	}

	if flag == true {
		tool.Failed(ctx, "手机号已被使用")
		return
	}

	tool.Success(ctx, "")
}

func (u *UserController) changePhone(ctx *gin.Context) {
	//获取并解析表单
	var phoneChangeParam param.ChangePhoneInfo
	err := ctx.ShouldBind(&phoneChangeParam)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
		return
	}

	//解析token
	gs := service.GeneralService{}
	us := service.UserService{}
	claims, err := gs.ParseToken(phoneChangeParam.Token)

	if err != nil {
		if err.Error()[:16] == "token is expired" {
			tool.Failed(ctx, "token失效")
			return
		}

		fmt.Println("getTokenParseTokenErr:", err)
		tool.Failed(ctx, "refreshToken不正确或系统错误")
		return
	}

	userinfo := claims.Userinfo

	//判断原设备类型
	if strings.Index(phoneChangeParam.OriginalAddress, "@") == -1 {
		//原设备为手机号
		if phoneChangeParam.OriginalAddress != userinfo.Phone {
			tool.Failed(ctx, "请输入原先绑定的手机号")
			return
		}
	} else {
		//原设备为email
		if phoneChangeParam.OriginalAddress != userinfo.Email {
			tool.Failed(ctx, "请输入原先绑定的email")
			return
		}
	}

	if phoneChangeParam.OriginalCode == "" {
		tool.Failed(ctx, "请输入验证码")
		return
	}

	//验证原设备
	flag, err := us.JudgeVerifyCode(ctx, phoneChangeParam.OriginalAddress, phoneChangeParam.OriginalCode)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("JudgeVerifyCode: ", err)
		return
	}

	if flag == false {
		tool.Failed(ctx, "原验证码错误")
		return
	}

	//验证新设备
	flag, err = us.JudgePhone(phoneChangeParam.NewPhone)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("JudgeEmailErr: ", err)
		return
	}

	if flag == true {
		tool.Failed(ctx, "该手机号已被使用")
		return
	}

	flag, err = us.JudgeVerifyCode(ctx, phoneChangeParam.NewPhone, phoneChangeParam.NewCode)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("JudgeVerifyCode: ", err)
		return
	}

	if flag == false {
		tool.Failed(ctx, "新手机号验证码错误")
		return
	}

	err = us.ChangePhone(userinfo.Username, phoneChangeParam.NewPhone)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("ChangeEmailErr: ", err)
		return
	}

	tool.Success(ctx, "")
}

func (u *UserController) changeStatement(ctx *gin.Context) {
	token := ctx.PostForm("token")
	newStatement := ctx.PostForm("new_statement")

	if newStatement == "" {
		newStatement = "这个人很懒，什么都没有写"
	}

	if len(newStatement) > 90 {
		tool.Failed(ctx, "这个个性签名太长了")
		return
	}

	gs := service.GeneralService{}
	us := service.UserService{}
	claims, err := gs.ParseToken(token)

	if err != nil {
		if err.Error()[:16] == "token is expired" {
			tool.Failed(ctx, "token失效")
			return
		}

		fmt.Println("getTokenParseTokenErr:", err)
		tool.Failed(ctx, "refreshToken不正确或系统错误")
		return
	}

	username := claims.Userinfo.Username
	err = us.ChangeStatement(username, newStatement)
	if err != nil {
		fmt.Println("ChangeStatementErr: ", err)
		tool.Failed(ctx, "系统错误")
	}

	tool.Success(ctx, "")
}

func (u *UserController) sendEmailCode(ctx *gin.Context) {
	email := ctx.PostForm("email")

	us := service.UserService{}
	verifyCode, err := us.SendCodeByEmail(email)

	if err != nil {
		if err.Error()[:3] == "cod" {
			tool.Failed(ctx, "邮箱不正确")
			return
		}
		fmt.Println("SendCodeByEmailErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	//把验证码放到redis中
	err = us.VerifyCodeIn(ctx, email, verifyCode)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("SetRedisErr: ", err)
		return
	}

	tool.Success(ctx, "")

}

func (u *UserController) getSelfInfo(ctx *gin.Context) {
	token := ctx.Query("token")

	gs := service.GeneralService{}
	claims, err := gs.ParseToken(token)

	if err != nil {
		if err.Error()[:16] == "token is expired" {
			tool.Failed(ctx, "token失效")
			return
		}

		fmt.Println("getTokenParseTokenErr:", err)
		tool.Failed(ctx, "refreshToken不正确或系统错误")
		return
	}

	userMap := tool.ObjToMap(claims.Userinfo)
	ctx.JSON(200, userMap)
}

func (u *UserController) changeEmail(ctx *gin.Context) {
	//获取并解析表单
	var emailChangeParam param.ChangeEmailInfo
	err := ctx.ShouldBind(&emailChangeParam)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
		return
	}

	//解析token
	gs := service.GeneralService{}
	us := service.UserService{}
	claims, err := gs.ParseToken(emailChangeParam.Token)

	if err != nil {
		if err.Error()[:16] == "token is expired" {
			tool.Failed(ctx, "token失效")
			return
		}

		fmt.Println("getTokenParseTokenErr:", err)
		tool.Failed(ctx, "refreshToken不正确或系统错误")
		return
	}

	userinfo := claims.Userinfo

	//判断原设备类型
	if strings.Index(emailChangeParam.OriginalAddress, "@") == -1 {
		//原设备为手机号
		if emailChangeParam.OriginalAddress != userinfo.Phone {
			tool.Failed(ctx, "请输入原先绑定的手机号")
			return
		}
	} else {
		//原设备为email
		if emailChangeParam.OriginalAddress != userinfo.Email {
			tool.Failed(ctx, "请输入原先绑定的email")
			return
		}
	}

	if emailChangeParam.OriginalCode == "" {
		tool.Failed(ctx, "请输入验证码")
		return
	}

	//验证原设备
	flag, err := us.JudgeVerifyCode(ctx, emailChangeParam.OriginalAddress, emailChangeParam.OriginalCode)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("JudgeVerifyCode: ", err)
		return
	}

	if flag == false {
		tool.Failed(ctx, "原验证码错误")
		return
	}

	//验证新设备
	flag, err = us.JudgeEmail(emailChangeParam.NewEmail)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("JudgeEmailErr: ", err)
		return
	}

	if flag == true {
		tool.Failed(ctx, "该邮箱已被使用")
		return
	}

	flag, err = us.JudgeVerifyCode(ctx, emailChangeParam.NewEmail, emailChangeParam.NewCode)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("JudgeVerifyCode: ", err)
		return
	}

	if flag == false {
		tool.Failed(ctx, "新邮箱验证码错误")
		return
	}

	err = us.ChangeEmail(userinfo.Username, emailChangeParam.NewEmail)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("ChangeEmailErr: ", err)
		return
	}

	tool.Success(ctx, "")
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
		"status":       "0",
		"data":         "登录成功",
		"token":        tokenString,
		"refreshToken": refreshToken,
	})
}

//检验用户名是否重复，是否符合规范
func (u *UserController) judgeUsername(ctx *gin.Context) {
	username := ctx.Query("username")

	if username == "" {
		tool.Failed(ctx, "请告诉我你的昵称吧")
		return
	}

	if len(username) > 14 {
		tool.Failed(ctx, "昵称过长")
		return
	}

	us := service.UserService{}
	flag, err := us.JudgeUsername(username)
	if err != nil {
		tool.Failed(ctx, "服务器出错")
		fmt.Println("服务器出错")
		return
	}

	if flag == false {
		tool.Success(ctx, "")
		return
	}

	tool.Failed(ctx, "昵称已存在")
}

//发送短信验证码
func (u *UserController) sendSms(ctx *gin.Context) {
	phone := ctx.PostForm("phone")

	if phone == "" {
		tool.Failed(ctx, "手机号不可为空")
	}

	us := service.UserService{}
	verifyCode, err := us.SendCodeByPhone(phone)
	if err != nil {
		tool.Failed(ctx, "系统错误")
		fmt.Println("SendCodeByPhoneErr")
		return
	}

	if verifyCode == "isv.MOBILE_NUMBER_ILLEGAL" {
		tool.Failed(ctx, "手机号不合法")
		fmt.Println("sendCodeByPhoneErr")
		return
	}

	//把验证码放到redis中
	err = us.VerifyCodeIn(ctx, phone, verifyCode)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("SetRedisErr: ", err)
		return
	}

	tool.Success(ctx, "")
}

func (u *UserController) register(ctx *gin.Context) {
	//获取并解析用户表单
	var userParam param.UserParam
	err := ctx.ShouldBind(&userParam)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
		return
	}

	//判断用户名
	if len(userParam.Username) > 15 {
		tool.Failed(ctx, "用户名太长了")
		return
	}

	if len(userParam.Username) == 0 {
		tool.Failed(ctx, "用户名不能为空")
		return
	}

	//判断密码是否正确
	if len(userParam.Pwd) < 6 {
		tool.Failed(ctx, " 密码不能小于6个字符")
		return
	}

	if len(userParam.Pwd) > 16 {
		tool.Failed(ctx, " 密码不能大于16个字符")
		return
	}

	//判断手机号是否可以使用
	phone := userParam.Phone

	us := service.UserService{}
	flag, err := us.JudgePhone(phone)
	if err != nil {
		fmt.Println("JudgePhoneErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if flag == true {
		tool.Failed(ctx, "该电话已经被注册")
		return
	}

	//判断验证码是否正确
	givenCode := userParam.VerifyCode

	if givenCode == "" {
		tool.Failed(ctx, "请输入验证码")
		return
	}

	flag, err = us.JudgeVerifyCode(ctx, phone, givenCode)
	if err != nil {
		if err.Error() == "redis: nil" {
			tool.Failed(ctx, "未发送验证码")
			return
		}
		tool.Failed(ctx, "服务器错误")
		fmt.Println("JudgePhoneCodeErr: ", err)
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
	err = us.RegisterModelIn(user)
	if err != nil {
		fmt.Println("RegisterModelInErr: ", err)
		tool.Failed(ctx, "服务器错误")
	}

	tool.Success(ctx, "")
}
