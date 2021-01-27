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
	"strings"
	"time"
)

type UserService struct {
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
func (u *UserService) JudgeUsername(username string) bool {
	d := dao.UserDao{tool.GetDb()}
	_, err := d.QueryByUsername(username)


	if err != nil && err.Error() == "sql: no rows in result set"{
		return false
	}

	return true
}

//检验手机是否存在, false不存在 反之存在
func (u *UserService) JudgePhone(phone string) bool {
	d := dao.UserDao{tool.GetDb()}
	_, err := d.QueryByPhone(phone)

	if err != nil && err.Error() == "sql: no rows in result set" {
		return false
	}

	return true
}

//检验验证码是否正确
func (u *UserService) JudgePhoneCode(ctx *gin.Context, key string, givenValue string) (bool, error) {
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
func (u *UserService) PhoneCodeIn(ctx *gin.Context, key string, value string) error {
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

	return "", nil
}
