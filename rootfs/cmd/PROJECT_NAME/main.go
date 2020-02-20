package main

import (
	"fmt"

	_ "<PROJECT_NAME>/config"
	"<PROJECT_NAME>/models"
	"<PROJECT_NAME>/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if viper.GetBool("mysql.enable") {
		log.Infof("Contect to mysql...")
		dsn := viper.GetString("mysql.dsn")
		models.InitMysql(dsn)
	} else {
		log.Warn("Mysql is disabled.")
	}

	addr := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	tls := viper.GetBool("tls.enable")
	certFile := viper.GetString("tls.certFile")
	keyFile := viper.GetString("tls.keyFile")
	server.StartServer(addr, tls, certFile, keyFile)
}
