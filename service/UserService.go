package service

import (
	"bilibili/dao"
	"bilibili/tool"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type UserService struct {

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
func (u *UserService) RegisterModelIn()  {
	
}

//通过手机号发送验证码
func (u *UserService) SendCodeByPhone(phone string) (string, error) {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	//调用阿里云sdk
	cfg := tool.GetCfg().Sms
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
