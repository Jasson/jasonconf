package conf

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	// Cfg .
	Cfg *ini.File
)

// 加载配置文件
func Init(path string) {
	log.Println("path:=", path)
	var err error

	if path == "" {
		Cfg, err = ini.Load("conf/app.ini")
	} else {
		Cfg, err = ini.Load(path)
	}
	if err != nil {
		log.Fatal("Load config err:=", err, "path:=", path)
	}
	log.Println("cfg:=", Cfg)

}

func GetConf(secStr, key string) string {
	sec, _ := Cfg.GetSection(secStr)
	return sec.Key(key).String()
}

func GetConfReturnInt(secStr, key string) int64 {

	log.Println("secStr=",secStr, "\n key:=", key)
	sec, _ := Cfg.GetSection(secStr)
	ret, err := sec.Key(key).Int64()
	if err != nil {
		log.Println("err:=", err)
		return 0
	}
	return ret
}
