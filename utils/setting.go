package utils

import (
	"gopkg.in/ini.v1"
	_ "gopkg.in/ini.v1"
)

var Conf = new(Config)

type Config struct {
	AppName  string `ini:"app_name"`
	Mode     string `ini:"mode"`
	HTTPPort string `ini:"http_port"`

	MySQL MySQLConfig `ini:"mysql"`

	TOCR tencentOCR `ini:"tencentOCR"`

	TROCR trWebOCR `ini:"TrWebOCR"`
}

type MySQLConfig struct {
	IP       string `ini:"ip"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type tencentOCR struct {
	SecretID  string `ini:"secretID"`
	SecretKey string `ini:"secretKey"`
	EndPoint  string `ini:"endPoint"`
	Region    string `ini:"region"`
}

type trWebOCR struct {
	Url string `ini:"url"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}


