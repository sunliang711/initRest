package router

import (
	"time"

	"<PROJECT_NAME>/handlers"
	"<PROJECT_NAME>/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// StartServer starts gin server
func StartServer(addr string, tls bool, certFile string, keyFile string) {
	//MUST SetMode first
	switch viper.GetString("server.gin_mode") {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	corsCfg := cors.Config{
		AllowOrigins: viper.GetStringSlice("cors.allow_origins"),
		AllowMethods: viper.GetStringSlice("cors.allow_methods"),
		AllowHeaders: viper.GetStringSlice("cors.allow_headers"),
		MaxAge:       time.Second * time.Duration(viper.GetInt("cors.max_age")),
	}
	logrus.Infof(utils.CorsConfigStringify(&corsCfg))

	router.Use(cors.New(corsCfg))

	// Put normal handlers below
	router.GET("/api/v1/health", handlers.Health)
	// router.GET("/api/v1/PATH", handlers.XXX)

	// Put need-auth handlers below
	// router.GET("/api/v1/PATH", middleware.Auth)
	// router.POST("/api/v1/PATH", middleware.Auth)


	logrus.Infof("Start server on %v, tls enabled: %v", addr, tls)
	if tls {
		logrus.Fatalln(router.RunTLS(addr, certFile, keyFile))
	} else {
		logrus.Fatalln(router.Run(addr))
	}

}
