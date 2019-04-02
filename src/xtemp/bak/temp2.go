package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name  string
	Age   int
	City  string
	Email string
	email string
	Pw    string
}

func main() {

	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// collection := client.Database("myweb").Collection("temp")

	/* ash := Trainer{"Ash", 10, "Pallet Town"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID) */

	/* misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}
	trainers := []interface{}{misty, brock}
	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	*/

	/* filter := bson.D{{"name", "Ash"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	*/

	collection := client.Database("myweb").Collection("user")
	// filter := bson.D{{"name", "Ash"}}
	filter := bson.D{{"email", "zhyuzh3d@hotmail.com"}}
	// filter := bson.M{"Email": "zhyuzh3d44@hotmail.com"}
	// filter := Trainer{Email: "zhyuzh3d44@hotmail.com"}
	// filter := Trainer{email: "zhyuzh3d@hotmail.com"}
	// filter := bson.M{"name": "Ash"}
	var result Trainer
	err4 := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err4 != nil {
		log.Fatal(err4)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	fmt.Println(collection.FindOne(context.TODO(), filter))

}
