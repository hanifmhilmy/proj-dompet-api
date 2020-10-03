package config

import (
	"log"

	"github.com/go-gcfg/gcfg"
)

// Config is struct to store all the config for the app
type Config struct {
	Database map[string]*struct {
		Master        string `gcfg:"master"`
		MasterMaxConn int    `gcfg:"master-conn"`
		MasterMaxIdle int    `gcfg:"master-idle"`
		MaxLifeConn   int64  `gcfg:"maxlifeconn"`
	}
}

//InitConfig public function to initialize the config
func InitConfig() (cnf Config, err error) {
	err = gcfg.ReadFileInto(&cnf, "config/db.main.ini")
	if err != nil {
		log.Println("Fail to read config file")
	}

	return
}
