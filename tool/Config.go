package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppPort  string         `json:"app_port"`
	AppHost  string         `json:"app_host"`
	Database DatabaseConfig `json:"database"`
	Email    EmailConfig    `json:"email"`
	Redis    RedisConfig    `json:"redis"`
	Jwt      JwtCfg         `json:"jwt"`
	Sms		 SmsCfg 		`json:"sms"`
}

type SmsCfg struct {
	SignName     string
	TemplateCode string
	AppKey       string
	AppSecret    string
	RegionId     string
}

type JwtCfg struct {
	SigningKey string `json:"signing_key"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type EmailConfig struct {
	ServiceEmail string `json:"service_email"`
	ServicePwd   string `json:"service_pwd"`
	SmtpPort     string `json:"smtp_port"`
	SmtpHost     string `json:"smtp_host"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	ShowSql  bool   `json:"show_sql"`
	Charset  string `json:"charset"`
}

var cfg *Config

//获取全局配置文件
func GetCfg() *Config {
	return cfg
}

func init() {
	err := ParseCfg("./config/app.json")
	if err != nil {
		panic(err)
	}
}

func ParseCfg(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}
