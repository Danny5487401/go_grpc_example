package util

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 虽然mgo 十分好用且稳定, 但是由于mgo不再维护 不支持事务, 并且golang 推荐使用官方驱动 mongo driver. 所以更换成mongo driver.
var mgoCli *mongo.Client

func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}

func initEngine() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 连接:在本地的时候mgo 的mongodburl 可以写成127.0.0.1,但是mongo driver 必须写成 mongodb://127.0.0.1
	uri := "mongodb://ali.danny.games:27017"
	var err error
	mgoCli, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	if err := mgoCli.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")
}
