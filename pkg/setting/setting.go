package setting

import (
    "log"
    "time"

    "github.com/go-ini/ini"
)

//编写与配置项保持一致的结构体（App、Server、Database）
type App struct {
    JwtSecret string
    PageSize int
    RuntimeRootPath string

    PrefixUrl  string
    ImageSavePath string
    ImageMaxSize int
    ImageAllowExts []string
    ExportSavePath string
    LogSavePath string
    LogSaveName string
    LogFileExt string
    TimeFormat string
}

var AppSetting = &App{}

type Server struct {
    RunMode string
    HttpPort int
    ReadTimeout time.Duration
    WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
    Type string
    User string
    Password string
    Host string
    Name string
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

var RedisSetting = &Redis{}

func Setup() {
    Cfg, err := ini.Load("src/github.com/Henate/Bloggor/conf/app.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
    }

    err = Cfg.Section("app").MapTo(AppSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
    }

    AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

    err = Cfg.Section("server").MapTo(ServerSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
    }

    //特殊设置的配置项进行再赋值
    ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
    ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

    //使用 MapTo 将conf/app.ini中的配置项映射到结构体DatabaseSetting上
    err = Cfg.Section("database").MapTo(DatabaseSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
    }

    err = Cfg.Section("redis").MapTo(RedisSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
    }
}