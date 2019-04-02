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


// type Trainer struct {
//     Name string
//     Age  int
//     City string
// }

func main(){

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// defer cancel()

	// collection := client.Database("myweb").Collection("user")

	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	fmt.Println("xxxxx")

}





// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	defer cancel()

// 	type usr struct {
// 		Email string
// 		Pw    string
// 	}

// 	// bu := bson.M{"Email": "ds.Email", "Pw": "ds.Pw"}
// 	dbc := client.Database("myweb").Collection("user")

// 	var result usr
// 	filter := bson.M{"Email": "ds.Email"}
// 	err := dbc.FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Found a single document: %+v\n", result)

// 	a := dbc.FindOne(context.TODO(), filter)
// 	a.Decode(&result)
// 	fmt.Println(a, result)
// 	if a == nil {
// 		fmt.Println("xx")
// 	}

	// filter := bson.M{"Email": "zhyuzh3d@hotmail.com"}
	// var u usr
	// dbc.FindOne(context.TODO(), bu).Decode(&u)
	// fmt.Println(u)

	// dbc.FindId(bson.ObjectIdHex("5c9ce6d692c49d2020b2b2e6")).One(&u)
	// fmt.Println(u)

	// dbc.InsertOne(context.TODO(), bu)

}

// type a struct {
// 	Val string
// 	b   struct {
// 		Address string
// 		Port    string
// 	}
// }

// func main() {

// 	c := &a{
// 		Val: "test",
// 		b: {
// 			Address: "addr",
// 			Port:    "port",
// 		},
// 	}
// 	fmt.Println(c)
// }

// func main() {
// 	// a := testmap()
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	defer cancel()

// 	fmt.Println(reflect.TypeOf(client))
// }

// type person struct {
// 	Name string
// 	Age  int
// }

// func testmap() map[string]interface{} {

// 	a := make(map[string]interface{})
// 	a["xx"] = person{Name: "lala", Age: 333}
// 	return a
// }

// func main() {
// 	mailRe, _ := regexp.Compile(`^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
// 	pwRe, _ := regexp.Compile(`^[0-9a-zA-Z_@]{6,18}$`)
// 	res := pwRe.MatchString("xxxx")
// 	res2 := mailRe.MatchString("xxxx@33.com")
// 	fmt.Println(res, res2)
// }
