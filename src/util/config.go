package util

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	MySql *MySqlConfig
}

func NewConfig() *Config {
	return &Config{
		MySql: NewMySqlConfig(),
	}
}

type MySqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewMySqlConfig() *MySqlConfig {
	str := os.Getenv("MYSQL_SETTING")
	var mysqlConfig MySqlConfig
	if err := json.Unmarshal([]byte(str), &mysqlConfig); err != nil {
		log.Panicf("NewMySqlConfig failed %s", err)
	}
	return &mysqlConfig
}
