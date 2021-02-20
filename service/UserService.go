package service

import (
	"bilibili/dao"
	"bilibili/model"
	"bilibili/tool"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type UserService struct {
}

func (u *UserService) GetUidByEmail(email string) (int64, error) {
	d := dao.UserDao{tool.GetDb()}

	userinfo, err := d.QueryByEmail(email)
	if err != nil {
		return 0, err
	}

	return userinfo.Uid, nil
}

func (u *UserService) GetUidByPhone(phone string) (int64, error) {
	d := dao.UserDao{tool.GetDb()}

	userinfo, err := d.QueryByPhone(phone)
	if err != nil {
		return 0, err
	}

	return userinfo.Uid, nil
}

func (u *UserService) ChangePassword(uid int64, newPassword string) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdatePassword(uid, newPassword)
	return err
}

func (u *UserService) ChangeUsername(uid int64, newUsername string) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdateUsername(uid, newUsername)
	if err != nil {
		return err
	}

	err = d.UpdateCoins(uid, -6)
	return err
}

func (u *UserService) GetUserinfo(uid int64) (model.Userinfo, error) {
	d := dao.UserDao{tool.GetDb()}

	userinfo, err := d.QueryByUid(uid)
	return userinfo, err
}

//签到服务
func (u *UserService) CheckIn(uid int64) error {
	d := dao.UserDao{tool.GetDb()}

	//加经验
	err := d.UpdateExp(uid, 5)
	if err != nil {
		return err
	}
	//加硬币
	err = d.UpdateCoins(uid, 1)
	if err != nil {
		return err
	}
	//更新日期
	err = d.UpdateLastCheckInDate(uid)
	return err
}

//可以签到返回true，否则返回false
func (u *UserService) JudgeCheckIn(uid int64) (bool, error) {
	d := dao.UserDao{tool.GetDb()}

	userinfo, err := d.QueryByUid(uid)
	if err != nil {
		return false, err
	}

	lastCheckInDate := userinfo.LastCheckInDate[:10]
	timeNow := time.Now().Format("2006-01-02")

	if timeNow == lastCheckInDate {
		return false, nil
	}

	return true, nil
}

func (u *UserService) ChangeStatement(uid int64, newStatement string) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdateStatement(uid, newStatement)
	return err
}

func (u *UserService) SendCodeByEmail(email string) (string, error) {
	email = strings.ToLower(email)
	emailCfg := tool.GetCfg().Email

	auth := smtp.PlainAuth("", emailCfg.ServiceEmail, emailCfg.ServicePwd, emailCfg.SmtpHost)
	to := []string{email}

	fmt.Println("EMAIL", email)

	rand.Seed(time.Now().Unix())
	code := rand.Intn(1000000)
	str := fmt.Sprintf("From:%v\r\nTo:%v\r\nSubject:bilibili验证码\r\n\r\n您的验证码为：%d\r\n请在10分钟内完成验证", emailCfg.ServiceEmail, email, code)
	msg := []byte(str)
	err := smtp.SendMail(emailCfg.SmtpHost+":"+emailCfg.SmtpPort, auth, emailCfg.ServiceEmail, to, msg)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(code), nil
}

func (u *UserService) ChangeBirthday(uid int64, newBirth time.Time) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdateBirthday(uid, newBirth)
	return err
}

func (u *UserService) ChangeGender(uid int64, newGender string) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdateGender(uid, newGender)
	return err
}

func (u *UserService) ChangeAvatar(uid int64, url string) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdateAvatar(uid, url)
	return err
}

func (u *UserService) ChangePhone(uid int64, newEmail string) error {
	d := dao.UserDao{tool.GetDb()}
	newEmail = strings.ToLower(newEmail)

	err := d.UpdatePhone(uid, newEmail)
	return err
}

func (u *UserService) ChangeEmail(uid int64, newEmail string) error {
	d := dao.UserDao{tool.GetDb()}
	newEmail = strings.ToLower(newEmail)

	err := d.UpdateEmail(uid, newEmail)
	return err
}

//通过短信登录
func (u *UserService) LoginBySms(phone string) (model.Userinfo, error) {
	d := dao.UserDao{tool.GetDb()}

	userinfo, err := d.QueryByPhone(phone)
	return userinfo, err
}

//通过密码登录，返回一个实体
func (u *UserService) Login(loginName, password string) (model.Userinfo, bool, error) {
	d := dao.UserDao{tool.GetDb()}

	//判断登录类型
	flag := strings.Index(loginName, "@")
	if flag != -1 {
		//邮箱登录
		loginName = strings.ToLower(loginName)
		userinfo, err := d.QueryByEmail(loginName)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return model.Userinfo{}, false, nil
			}
			return model.Userinfo{}, false, err
		}

		//md5解密
		m5 := md5.New()
		m5.Write([]byte(password))
		m5.Write([]byte(userinfo.Salt))
		st := m5.Sum(nil)
		hashPwd := hex.EncodeToString(st)

		if hashPwd != userinfo.Password {
			return model.Userinfo{}, false, nil
		}
		return userinfo, true, nil
	} else {
		//手机号登录

		userinfo, err := d.QueryByPhone(loginName)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return model.Userinfo{}, false, nil
			}
			return model.Userinfo{}, false, err
		}

		//md5解密
		m5 := md5.New()
		m5.Write([]byte(password))
		m5.Write([]byte(userinfo.Salt))
		st := m5.Sum(nil)
		hashPwd := hex.EncodeToString(st)

		if hashPwd != userinfo.Password {
			return model.Userinfo{}, false, nil
		}

		return userinfo, true, nil
	}
}

//检验用户名是否存在, false不存在 反之存在
func (u *UserService) JudgeUsername(username string) (bool, error) {
	d := dao.UserDao{tool.GetDb()}
	_, err := d.QueryByUsername(username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

//检验手机是否存在, false不存在 反之存在
func (u *UserService) JudgePhone(phone string) (bool, error) {
	d := dao.UserDao{tool.GetDb()}
	_, err := d.QueryByPhone(phone)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

//检验邮箱是否存在, false不存在 反之存在
func (u *UserService) JudgeEmail(email string) (bool, error) {
	d := dao.UserDao{tool.GetDb()}
	email = strings.ToLower(email)
	_, err := d.QueryByEmail(email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

//检验验证码是否正确
func (u *UserService) JudgeVerifyCode(ctx *gin.Context, key string, givenValue string) (bool, error) {
	rd := dao.RedisDao{}
	value, err := rd.RedisGetValue(ctx, key)
	if err != nil {
		return false, err
	}

	if value != givenValue {
		return false, nil
	}

	return true, nil

}

//验证码放入redis中
func (u *UserService) VerifyCodeIn(ctx *gin.Context, key string, value string) error {
	rd := dao.RedisDao{}
	err := rd.RedisSetValue(ctx, key, value)
	return err
}

//注册实体放入mysql
func (u *UserService) RegisterModelIn(userinfo model.Userinfo) error {
	d := dao.UserDao{tool.GetDb()}
	err := d.InsertUser(userinfo)
	return err
}

//通过手机号发送验证码
func (u *UserService) SendCodeByPhone(phone string) (string, error) {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	//调用阿里云sdk
	cfg := tool.GetCfg().Sms
	//fmt.Println("asdfsadf", cfg.AppSecret, cfg.AppKey)
	client, err := dysmsapi.NewClientWithAccessKey(cfg.RegionId, cfg.AppKey, cfg.AppSecret)
	if err != nil {
		return "", err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = cfg.SignName
	request.TemplateCode = cfg.TemplateCode
	request.PhoneNumbers = phone

	par, err := json.Marshal(gin.H{
		"code": code,
	})

	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	fmt.Println(response)

	if err != nil {
		return "", err
	}

	//成功
	if response.Code == "OK" {
		return code, nil
	}

	if response.Code == "isv.MOBILE_NUMBER_ILLEGAL" {
		return "isv.MOBILE_NUMBER_ILLEGAL", nil
	}

	return "", nil
}
