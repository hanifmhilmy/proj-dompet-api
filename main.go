package main

import (
	"log"

	"github.com/hanifmhilmy/proj-dompet-api/app/config"
)

var cnf config.Config

func init() {
	cnf, err := config.InitConfig()
	if err != nil {
		log.Panic("[Failed to initialize] Config is not set properly please check!")
	}
}
