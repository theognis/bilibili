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
	engine.POST("/api/verify/sms/login", u.sendSmsLogin)
	engine.POST("/api/user/login/pw", u.login)
	engine.POST("/api/user/login/sms", u.loginBySms)
	engine.POST("/api/verify/email", u.sendEmailCode)
	engine.PUT("/api/user/username", u.changeUsername)
	engine.PUT("/api/user/password", u.changePassword)
	engine.PUT("/api/user/phone", u.changePhone)
	engine.PUT("/api/user/email", u.changeEmail)
	engine.PUT("/api/user/statement", u.changeStatement)
	engine.PUT("/api/user/check-in", u.checkIn)
	engine.PUT("/api/user/avatar", u.changeAvatar)
	engine.PUT("/api/user/gender", u.changeGender)
	engine.PUT("/api/user/birth", u.changeBirth)
}

func (u *UserController) changePassword(ctx *gin.Context) {
	var ChangePasswordParam param.ChangePasswordParam
	err := ctx.BindJSON(&ChangePasswordParam)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
		return
	}

	token := ChangePasswordParam.Token
	if token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	us := service.UserService{}
	gs := service.TokenService{}
	//解析token
	clams, err := gs.ParseToken(token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo

	//检测账号相关
	if ChangePasswordParam.Account == "" {
		tool.Failed(ctx, "账号为空")
		return
	}

	if strings.Index(ChangePasswordParam.Account, "@") == -1 {
		//手机号
		flag, err = us.JudgePhone(ChangePasswordParam.Account)
		if err != nil {
			fmt.Println("JudgePhoneErr: ", err)
			tool.Failed(ctx, "服务器错误")
			return
		}

		if flag == false {
			tool.Failed(ctx, "账号不存在")
			return
		}
	} else {
		//邮箱
		flag, err = us.JudgeEmail(ChangePasswordParam.Account)
		if err != nil {
			fmt.Println("JudgeEmailErr: ", err)
			tool.Failed(ctx, "服务器错误")
			return
		}

		if flag == false {
			tool.Failed(ctx, "账号不存在")
		}
	}

	//验证码相关
	if ChangePasswordParam.Code == "" {
		tool.Failed(ctx, "验证码为空")
		return
	}

	flag, err = us.JudgeVerifyCode(ctx, ChangePasswordParam.Account, ChangePasswordParam.Code)
	if err != nil {
		fmt.Println("JudgeVerifyCodeErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if flag == false {
		tool.Failed(ctx, "验证码错误")
		return
	}

	//验证新密码
	if len(ChangePasswordParam.NewPassword) < 6 {
		tool.Failed(ctx, "密码不能小于6个字符")
		return
	}

	if len(ChangePasswordParam.NewPassword) > 16 {
		tool.Failed(ctx, "密码不能大于16个字符")
		return
	}

	uid := userinfo.Uid
	err = us.ChangePassword(uid, ChangePasswordParam.NewPassword)
	if err != nil {
		fmt.Println("ChangePasswordErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	tool.Success(ctx, "")
}

func (u *UserController) changeUsername(ctx *gin.Context) {
	token := ctx.PostForm("token")
	newUsername := ctx.PostForm("new_username")

	if token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	us := service.UserService{}
	gs := service.TokenService{}
	//解析token
	clams, err := gs.ParseToken(token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo

	if newUsername == "" {
		tool.Failed(ctx, "昵称不可为空")
		return
	}

	if len(newUsername) > 15 {
		tool.Failed(ctx, "昵称太长了")
		return
	}

	if userinfo.Coins < 6 {
		tool.Failed(ctx, "硬币不足")
		return
	}

	err = us.ChangeUsername(userinfo.Uid, newUsername)
	if err != nil {
		fmt.Println("ChangeUsernameErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	tool.Success(ctx, "")
}

//更改生日
func (u *UserController) changeBirth(ctx *gin.Context) {
	token := ctx.PostForm("token")
	newBirth := ctx.PostForm("new_birth")

	if token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	us := service.UserService{}
	gs := service.TokenService{}
	//解析token
	clams, err := gs.ParseToken(token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo

	newTime, err := time.ParseInLocation("2006-01-02", newBirth, time.Local)
	if err != nil {
		fmt.Println("ParseInLocationErr: ", err)
		tool.Failed(ctx, "日期格式错误")
		return
	}

	if time.Now().Before(newTime) {
		tool.Failed(ctx, "出生日期无效")
		return
	}

	err = us.ChangeBirthday(userinfo.Uid, newTime)
	if err != nil {
		fmt.Println("ChangeBirthdayErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	tool.Success(ctx, "")
}

//更改性别
func (u *UserController) changeGender(ctx *gin.Context) {
	token := ctx.PostForm("token")
	newGender := ctx.PostForm("new_gender")

	if token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	us := service.UserService{}
	gs := service.TokenService{}
	//解析token
	clams, err := gs.ParseToken(token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo

	if newGender != "F" && newGender != "M" && newGender != "O" && newGender != "N" {
		tool.Failed(ctx, "无效的性别")
		return
	}

	err = us.ChangeGender(userinfo.Uid, newGender)
	if err != nil {
		tool.Failed(ctx, "系统错误")
		fmt.Println("ChangeGenderErr: ", err)
		return
	}

	tool.Success(ctx, "")
}

//使用短信登录
func (u *UserController) loginBySms(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	verifyCode := ctx.PostForm("verify_code")

	if phone == "" {
		tool.Failed(ctx, "手机号不能为空哦")
		return
	}

	if verifyCode == "" {
		tool.Failed(ctx, "短信验证码不能为空")
		return
	}

	us := service.UserService{}
	flag, err := us.JudgeVerifyCode(ctx, phone, verifyCode)
	if err != nil {
		fmt.Println("JudgePhoneErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if flag == false {
		tool.Failed(ctx, "验证码错误")
		return
	}

	userinfo, err := us.LoginBySms(phone)
	if err != nil {
		fmt.Println("loginErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	gs := service.TokenService{}

	//创建token， 有效期两分钟
	tokenString, err := gs.CreateToken(userinfo, 120, "TOKEN")
	if err != nil {
		fmt.Println("CreateTokenErr:", err)
		tool.Failed(ctx, "系统错误")
		return
	}

	//创建一个refresh token有效期一周
	refreshToken, err := gs.CreateToken(userinfo, 604800, "REFRESH_TOKEN")
	if err != nil {
		fmt.Println("CreateRefreshTokenErr:", err)
		tool.Failed(ctx, "系统错误")
		return
	}

	ctx.JSON(200, gin.H{
		"status":       true,
		"data":         "",
		"token":        tokenString,
		"refreshToken": refreshToken,
	})
}

//更改用户头像
func (u *UserController) changeAvatar(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("avatar")
	token := ctx.PostForm("token")

	if err != nil {
		fmt.Println("FormFileErr: ", err)
		tool.Failed(ctx, "上传失败")
		return
	}

	if token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	us := service.UserService{}
	gs := service.TokenService{}
	//解析token
	clams, err := gs.ParseToken(token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo

	//大小限制2Mb
	if header.Size > (2 << 20) {
		tool.Failed(ctx, "头像文件过大")
		return
	}

	//格式限制
	extension := tool.GetExtension(header.Filename)
	extension = strings.ToLower(extension)
	if extension != "jpg" && extension != "png" {
		tool.Failed(ctx, "头像无效")
		return
	}

	Os := service.OssService{}

	fileName := strconv.FormatInt(userinfo.Uid, 10) + "." + extension
	err = Os.UploadAvatar(file, fileName)
	if err != nil {
		fmt.Println("UploadAvatarErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	cfg := tool.GetCfg().Oss
	url := cfg.AvatarUrl + fileName

	//头像入数据库
	err = us.ChangeAvatar(userinfo.Uid, url)
	if err != nil {
		fmt.Println("ChangeAvatarErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	tool.Success(ctx, "上传成功")
}

//签到
func (u *UserController) checkIn(ctx *gin.Context) {
	token := ctx.PostForm("token")

	flag := tool.CheckTokenNil(ctx, token)
	if flag == false {
		return
	}

	us := service.UserService{}
	gs := service.TokenService{}

	clams, err := gs.ParseToken(token)
	flag = tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}

	userinfo := clams.Userinfo

	//查询是否可以签到
	flag, err = us.JudgeCheckIn(userinfo.Uid)
	if err != nil {
		tool.Failed(ctx, "服务器错误")
		fmt.Println("JudgeCheckInErr: ", err)
		return
	}

	if flag == false {
		tool.Failed(ctx, "ALREADY_DONE")
		return
	}

	err = us.CheckIn(userinfo.Uid)
	if err != nil {
		tool.Failed(ctx, "系统错误")
		fmt.Println("CheckInErr: ", err)
		return
	}

	tool.Success(ctx, "SUCCESS")
}

func (u *UserController) sendSmsLogin(ctx *gin.Context) {
	phone := ctx.PostForm("phone")

	if phone == "" {
		tool.Failed(ctx, "手机号不可为空")
		return
	}

	us := service.UserService{}
	flag, err := us.JudgePhone(phone)
	if err != nil {
		fmt.Println("JudgePhoneErr: ", err)
		tool.Failed(ctx, "f服务器错误")
		return
	}

	if flag == false {
		tool.Failed(ctx, "手机号未被注册")
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

func (u *UserController) sendSmsRegister(ctx *gin.Context) {
	phone := ctx.PostForm("phone")

	if phone == "" {
		tool.Failed(ctx, "手机号不可为空")
		return
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

	if phoneChangeParam.Token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	us := service.UserService{}
	gs := service.TokenService{}
	//解析token
	clams, err := gs.ParseToken(phoneChangeParam.Token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo
	//判断原设备类型
	if strings.Index(phoneChangeParam.OldAccount, "@") == -1 {
		//原设备为手机号
		if phoneChangeParam.OldAccount != userinfo.Phone {
			tool.Failed(ctx, "请输入原先绑定的手机号")
			return
		}
	} else {
		//原设备为email
		if phoneChangeParam.OldAccount != userinfo.Email {
			tool.Failed(ctx, "请输入原先绑定的email")
			return
		}
	}

	if phoneChangeParam.OldCode == "" {
		tool.Failed(ctx, "请输入验证码")
		return
	}

	//验证原设备
	flag, err = us.JudgeVerifyCode(ctx, phoneChangeParam.OldAccount, phoneChangeParam.OldCode)
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

	err = us.ChangePhone(userinfo.Uid, phoneChangeParam.NewPhone)
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

	if token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	us := service.UserService{}
	gs := service.TokenService{}

	clams, err := gs.ParseToken(token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo

	err = us.ChangeStatement(userinfo.Uid, newStatement)
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
	if token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	gs := service.TokenService{}
	us := service.UserService{}

	clams, err := gs.ParseToken(token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	uid := clams.Userinfo.Uid

	userinfo, err := us.GetUserinfo(uid)
	if err != nil {
		fmt.Println("GetUserinfoErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	userMap := tool.ObjToMap(userinfo)
	tool.Success(ctx, userMap)
}

func (u *UserController) changeEmail(ctx *gin.Context) {
	//获取并解析表单
	var emailChangeParam param.ChangeEmailInfo
	err := ctx.BindJSON(&emailChangeParam)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
		return
	}

	if emailChangeParam.Token == "" {
		tool.Failed(ctx, "NO_TOKEN_PROVIDED")
		return
	}

	//解析token
	us := service.UserService{}
	gs := service.TokenService{}

	clams, err := gs.ParseToken(emailChangeParam.Token)
	flag := tool.CheckTokenErr(ctx, clams, err)
	if flag == false {
		return
	}
	userinfo := clams.Userinfo

	//判断原设备类型
	if strings.Index(emailChangeParam.OldAccount, "@") == -1 {
		//原设备为手机号
		if emailChangeParam.OldAccount != userinfo.Phone {
			tool.Failed(ctx, "请输入原先绑定的手机号")
			return
		}
	} else {
		//原设备为email
		if emailChangeParam.OldAccount != userinfo.Email {
			tool.Failed(ctx, "请输入原先绑定的email")
			return
		}
	}

	if emailChangeParam.OldCode == "" {
		tool.Failed(ctx, "请输入验证码")
		return
	}

	//验证原设备
	flag, err = us.JudgeVerifyCode(ctx, emailChangeParam.OldAccount, emailChangeParam.OldCode)
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

	err = us.ChangeEmail(userinfo.Uid, emailChangeParam.NewEmail)
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
	gs := service.TokenService{}

	if loginName == "" {
		tool.Failed(ctx, "请输入注册时用的邮箱或者手机号呀")
		return
	}

	if password == "" {
		tool.Failed(ctx, "喵，你没输入密码么？")
		return
	}

	userinfo, ok, err := us.Login(loginName, password)
	if err != nil {
		fmt.Println("loginErr: ", err)
		tool.Failed(ctx, "服务器错误")
		return
	}

	if ok == false {
		tool.Failed(ctx, "用户名或密码错误")
		return
	}

	//创建token， 有效期两分钟
	tokenString, err := gs.CreateToken(userinfo, 120, "TOKEN")
	if err != nil {
		fmt.Println("CreateTokenErr:", err)
		tool.Failed(ctx, "系统错误")
		return
	}

	//创建一个refresh token有效期一周
	refreshToken, err := gs.CreateToken(userinfo, 604800, "REFRESH_TOKEN")
	if err != nil {
		fmt.Println("CreateRefreshTokenErr:", err)
		tool.Failed(ctx, "系统错误")
		return
	}

	ctx.JSON(200, gin.H{
		"status":       true,
		"data":         "",
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
