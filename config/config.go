package config

import (
	"log"

	"github.com/go-gcfg/gcfg"
	"github.com/joho/godotenv"
)

const (
	SecretConst        = "API_SECRET"
	SecretRefreshConst = "API_REFRESH_SECRET"
)

// Config is struct to store all the config for the app
type Config struct {
	Database map[string]*struct {
		Master        string `gcfg:"master"`
		MasterMaxConn int    `gcfg:"master-conn"`
		MasterMaxIdle int    `gcfg:"master-idle"`
		MaxLifeConn   int64  `gcfg:"maxlifeconn"`
	}
	Redis struct {
		Main string `gcfg:"main"`
	}
	RedisOptions struct {
		MaxActive      int
		MaxIdle        int
		Timeout        int64
		IdlePingPeriod int64
		PoolWaitMs     int64
	}
	Token struct {
		AccessExpire  int64 `gcfg:"access-exp"`  // minutes
		RefreshExpire int64 `gcfg:"refresh-exp"` // days
	}
}

//InitConfig public function to initialize the config
func InitConfig() (cnf Config, err error) {
	// Read the env file
	err = godotenv.Load("config/files/app.env")
	if err != nil {
		log.Println("Fail to read env file")
	}

	err = gcfg.ReadFileInto(&cnf, "config/files/dompet.main.ini")
	if err != nil {
		log.Println("Fail to read config file")
	}

	return
}
