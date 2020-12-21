package database

import (
	"context"
	"database/sql"

	// register mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/sunliang711/goutils/mongodb"
	umysql "github.com/sunliang711/goutils/mysql"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/spf13/viper"
)

var (
	MysqlConn    *sql.DB
	MysqlORMConn *gorm.DB
	MongoConn    *mongo.Client
)

//InitMysql open mysql with dsn
func InitMysql(dsn string) {
	logrus.Infof("Try to connect to mysql: '%v'", dsn)
	var err error
	if viper.GetBool("mysql.orm") {
		MysqlORMConn, err = gorm.Open("mysql", dsn)
		logrus.Infof("Use gorm driver...")
		if err != nil {
			panic(err)
		}

	} else {
		MysqlConn, err = umysql.New(dsn, viper.GetInt("mysql.maxIdleConns"), viper.GetInt("mysql.maxOpenConns"))
		if err != nil {
			panic(err)
		}
	}
	logrus.Infof("Connected to mysql: '%v'", dsn)
}

//CloseMysql close mysql connection
func CloseMysql() {
	logrus.Infoln("Close mysql.")
	MysqlConn.Close()
}

// InitMongo opens a mongodb connection
func InitMongo(url string) {
	logrus.Infof("Try to connect to mongodb: '%v'", url)
	var err error
	MongoConn, err = mongodb.New(url, 5)
	if err != nil {
		panic(err)
	}
	logrus.Infof("Connected to mongodb: '%v'", url)
}

func CloseMongo() {
	logrus.Infof("Close mongodb")
	MongoConn.Disconnect(context.Background())
}
