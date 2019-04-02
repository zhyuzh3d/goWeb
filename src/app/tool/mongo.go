package tool

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//MongoDBCLient mongo数据库客户端，init中自动初始化
var MongoDBCLient *mongo.Client

func init() {
	MongoDBCLient = initMongoDB()
}

//initMongoDB 初始化工具集
func initMongoDB() *mongo.Client {
	//连接服务
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer cancel()

	if err != nil {
		panic(err)
	}

	//检测连接是否成功
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	defer cancel()
	if err != nil {
		panic(err)
	}

	return client
}
