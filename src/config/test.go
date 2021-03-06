package config

import (
	"time"

	"github.com/globalsign/mgo"
)

var test = Config{
	IsDebug:   true,
	Port:      3000,
	SecretKey: []byte("JjceTG4Hzfg7csHv"),

	Mongo: &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "demotest",
		Username: "demotest",
		Password: "123",
		Timeout:  5 * time.Second,
	},
}
