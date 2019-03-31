package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("myweb").Collection("user")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	a := collection.FindOne(ctx, bson.M{"Email": "zhyuzh3d@hotmail.com"})
	fmt.Println(a)
	b := collection.FindOne(ctx, bson.M{"Email": "zhyuz2h3d@hotmail.com"})
	fmt.Println(b)

	// var result bson.M
	// err := collection.FindOne(ctx, bson.M{"Email": "zhyuzh3d@hotmail.com"}).Decode(&result)
	// if err != nil {
	// 	log.Fatal(">>>", err)
	// }
	// fmt.Println(`bson.M{"name": "a"}->`)
	// fmt.Println("Name", result, result["Email"], result["Pw"])

	/* ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	c, _ := collection.CountDocuments(ctx, bson.M{"Email": "zhyuzh3d2@hotmail.com"})
	fmt.Println("--->")
	fmt.Println("Count", c) */
}
