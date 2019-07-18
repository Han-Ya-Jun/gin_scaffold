package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	Level              string
	FileWriterOn       bool
	Path               string
	RotateLogPath      string
	WfLogPath          string
	RotateWfLogPath    string
	ConsoleWriterOn    bool
	ConsoleWriterColor bool
	TimeFormat         string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Mongo struct {
	Host     string
	Password string
	Source   string
	User     string
}

var RedisSetting = &Redis{}
var MongoSetting = &Mongo{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() error {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
		return err
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("mongo", MongoSetting)
	//AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
	return nil
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting e: %v", err)
	}
}
