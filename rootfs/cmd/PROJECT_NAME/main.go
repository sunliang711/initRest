package main

import (
	"fmt"

	_ "haha/config"
	"haha/models"
	"haha/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if viper.GetBool("mysql.enable") {
		dsn := viper.GetString("mysql.dsn")
		models.InitMysql(dsn)
	} else {
		log.Info("Mysql is disabled.")
	}

	if viper.GetBool("mongodb.enable") {
		dsn := viper.GetString("mongodb.url")
		models.InitMongo(dsn)
	} else {
		log.Info("Mongodb is disabled.")
	}

	addr := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	tls := viper.GetBool("tls.enable")
	certFile := viper.GetString("tls.certFile")
	keyFile := viper.GetString("tls.keyFile")
	server.StartServer(addr, tls, certFile, keyFile)
}
