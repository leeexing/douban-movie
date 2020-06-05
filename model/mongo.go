package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/leeexing/douban-movie/setting"
)

var (
	mongodb *mongo.Database

	cName string
)

func init() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("mongo.Connect.Error:", err)
	}

	sec, err := setting.Cfg.GetSection("mongo")
	if err != nil {
		log.Fatalf("Fail to get section 'mongo': %v", err)
	}

	dName := sec.Key("NAME").MustString("mongo")
	cName = sec.Key("COLLECTION").MustString("mongo")

	mongodb = client.Database(dName)

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("client.Ping.Error:", err)
	}
	log.Println("connect to Mongo")

	// filter := bson.D{{Key: "title", Value: "肖申克的救赎"}}
	// 或者使用这种方式，优选第一种方式
	// filter := bson.D{primitive.E{Key: "title", Value: "肖申克的救赎"}}
	// GetOneMovie(&filter)
	// updateOneMovie(&filter)
}

// SaveMovieToMongo 保存豆瓣电影数据到mongo数据库中
func SaveMovieToMongo(movie *DoubanMovie) {
	c := mongodb.Collection(cName)
	res, err := c.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal("movie.InsertOne.Error:", err)
	}
	fmt.Println("+++ Inserted documents: ", res.InsertedID)
}

// GetOneMovie 查看文档
func GetOneMovie(filter *bson.D) {
	c := mongodb.Collection(cName)
	var res DoubanMovie
	err := c.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		log.Fatal("Found Movie Error:", err)
	}
	log.Printf("Found a single document: %+v", res)
}

// 更新一个豆瓣电影的文档
func updateOneMovie(filter *bson.D) {
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "like", Value: 1},
		}},
	}
	c := mongodb.Collection(cName)
	uRes, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Update one document", uRes.ModifiedCount)
}
