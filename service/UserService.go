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

func (u *UserService) GetUserinfo(username string) (model.Userinfo, error) {
	d := dao.UserDao{tool.GetDb()}

	userinfo, err := d.QueryByUsername(username)
	return userinfo, err
}

//签到服务
func (u *UserService) CheckIn(username string) error {
	d := dao.UserDao{tool.GetDb()}

	//加经验
	err := d.UpdateExp(username, 5)
	if err != nil {
		return err
	}
	//加硬币
	err = d.UpdateCoins(username, 1)
	if err != nil {
		return err
	}
	//更新日期
	err = d.UpdateLastCheckInDate(username)
	return err
}

//可以签到返回true，否则返回false
func (u *UserService) JudgeCheckIn(username string) (bool, error) {
	d := dao.UserDao{tool.GetDb()}

	userinfo, err := d.QueryByUsername(username)
	if err != nil {
		return false, err
	}

	lastCheckInDate := userinfo.LastCheckInDate[:10]
	timeNow := time.Now().Format("2006-01-02")

	fmt.Println("SQL:", lastCheckInDate)
	fmt.Println("NOW:", timeNow)

	if timeNow == lastCheckInDate {
		return false, nil
	}

	return true, nil
}

func (u *UserService) ChangeStatement(username, newStatement string) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdateStatement(username, newStatement)
	return err
}

func (u *UserService) SendCodeByEmail(email string) (string, error) {
	emailCfg := tool.GetCfg().Email

	auth := smtp.PlainAuth("", emailCfg.ServiceEmail, emailCfg.ServicePwd, emailCfg.SmtpHost)
	to := []string{email}

	fmt.Println("EMAIL", email)

	rand.Seed(time.Now().Unix())
	code := rand.Intn(10000)
	str := fmt.Sprintf("From:%v\r\nTo:%v\r\nSubject:tieba注册验证码\r\n\r\n您的验证码为：%d\r\n请在10分钟内完成验证", emailCfg.ServiceEmail, email, code)
	msg := []byte(str)
	err := smtp.SendMail(emailCfg.SmtpHost+":"+emailCfg.SmtpPort, auth, emailCfg.ServiceEmail, to, msg)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(code), nil
}

func (u *UserService) ChangeAvatar(username, url string) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdateAvatar(username, url)
	return err
}

func (u *UserService) ChangePhone(username, newEmail string) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdatePhone(username, newEmail)
	return err
}

func (u *UserService) ChangeEmail(username, newEmail string) error {
	d := dao.UserDao{tool.GetDb()}

	err := d.UpdateEmail(username, newEmail)
	return err
}

//返回一个实体
func (u *UserService) Login(loginName, password string) (model.Userinfo, bool, error) {
	d := dao.UserDao{tool.GetDb()}

	//判断登录类型
	flag := strings.Index(loginName, "@")
	if flag != -1 {
		//邮箱登录
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
