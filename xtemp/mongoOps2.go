package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("myweb").Collection("temp3")

	//写入单个数据bson.M
	/* ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	d := bson.M{"Name": "a", "Value": `bson.M{"Name": "a", "Value": ""}`}
	res, _ := collection.InsertOne(ctx, d)
	id := res.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(id) */

	//写入单个数据bson.M
	/* ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	d := bson.M{"name": "a", "value": `bson.M{"name": "a", "Value"`}
	res, _ := collection.InsertOne(ctx, d)
	id := res.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(id) */

	/* ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	d := bson.M{"name": "a", "Value": `bson.M{"name": "a", "Value"`}
	res, _ := collection.InsertOne(ctx, d)
	id := res.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(id) */

	//写入单个数据bson.M
	/* type D struct {
		Name  string
		Value string
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	d := D{Name: "a", Value: `D{Name: "a", Value`}
	res, _ := collection.InsertOne(ctx, d)
	id := res.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(id) */

	//写入单个数据bson.M
	/* type D struct {
		name  string
		value string
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	d := D{name: "a", value: `D{name: "a", value`}
	res, _ := collection.InsertOne(ctx, d)
	id := res.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(id) */

	//struct
	/* type D struct {
		Name  string
		Value float64
	}
	d := D{"struct-Name", 3.14}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, _ := collection.InsertOne(ctx, d)
	id := res.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(id) */

	//读取单个数据
	/* 	var result struct {
		value string
		name  string
		Value string
		Name  string
	} */
	// filter := bson.M{"Name": "a"}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var result bson.M
	err := collection.FindOne(ctx, bson.M{"Name": "a"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(`bson.M{"name": "a"}->`)
	// fmt.Println("name", result.name, result.value)
	fmt.Println("Name", result["Name"], result["Value"])

	fmt.Println("--->")
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.M{"Name": "a"})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result, "-->", result["Value"])
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// count, err := collection.Find(ctx, bson.M{"Name": "a"}).Count()
	// fmt.Println("Count", cur.Count())

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	c, _ := collection.CountDocuments(ctx, bson.M{})
	fmt.Println("--->")
	fmt.Println("Count", c)
	//
}

// for i := 0; i < s.NumField(); i++ {
// 	f := s.Field(i)
// 	fmt.Printf("%d: %s %s = %v\n", i,
// 		typeOfT.Field(i).Name, f.Type(), f.Interface())
// }

// // fmt.Println("Count", reflect.TypeOf(count))
