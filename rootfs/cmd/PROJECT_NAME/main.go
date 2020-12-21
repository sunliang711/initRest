package main

import (
	"fmt"

	"github.com/sunliang711/goutils/config"
	"<PROJECT_NAME>/database"
	"<PROJECT_NAME>/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	err := config.InitConfigLogger()
	if err != nil{
		panic(err)
	}

	if viper.GetBool("mysql.enable") {
		dsn := viper.GetString("mysql.dsn")
		database.InitMysql(dsn)
	} else {
		log.Info("Mysql is disabled.")
	}

	if viper.GetBool("mongodb.enable") {
		dsn := viper.GetString("mongodb.url")
		database.InitMongo(dsn)
	} else {
		log.Info("Mongodb is disabled.")
	}

	addr := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	tls := viper.GetBool("tls.enable")
	certFile := viper.GetString("tls.certFile")
	keyFile := viper.GetString("tls.keyFile")
	router.StartServer(addr, tls, certFile, keyFile)
}
