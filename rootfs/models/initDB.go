package models

import (
	"context"
	"database/sql"
	"time"

	// register mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db  *sql.DB
	mdb *mongo.Client
)

//InitMysql open mysql with dsn
func InitMysql(dsn string) {
	logrus.Infof("Try to connect to mysql: '%v'", dsn)
	var err error
	if len(dsn) == 0 {
		logrus.Fatal("Mysql DSN is empty.")
	}
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		logrus.Fatalf("Open mysql error: %v", err)
	}
	err = db.Ping()
	if err != nil {
		logrus.Fatalf("Ping mysql error: %v", err)
	}

	logrus.Infof("Connected to mysql: '%v'", dsn)
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
}

//CloseMysql close mysql connection
func CloseMysql() {
	logrus.Infoln("Close mysql.")
	db.Close()
}

// InitMongo opens a mongodb connection
func InitMongo(url string) {
	logrus.Infof("Try to connect to mongodb: '%v'", url)
	var err error
	mdb, err = mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = mdb.Connect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = mdb.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Fatalf("Connect to mongodb: %v failed", url)
	}
	logrus.Infof("Connected to mongodb: '%v'", url)
	return
}
