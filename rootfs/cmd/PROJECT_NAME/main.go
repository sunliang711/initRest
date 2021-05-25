package main

import (
	"fmt"

	"<PROJECT_NAME>/database"
	"<PROJECT_NAME>/router"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.Infof("main()")
	defer func(){
		database.Release()
	}()

	addr := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	tls := viper.GetBool("tls.enable")
	certFile := viper.GetString("tls.cert_file")
	keyFile := viper.GetString("tls.key_file")
	router.StartServer(addr, tls, certFile, keyFile)
}
