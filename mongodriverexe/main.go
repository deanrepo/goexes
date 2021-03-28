package main

import (
	"context"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	cappedCollectionTest()
}

//action type constants
const (
	VIEW_PRODUCT = iota
	ADD_TO_CART
	CHECKOUT
	PURCHASE
)

// test for capped collections
func cappedCollectionTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	isCapped := true
	var limitSize int64 = 16384
	err = client.Database("garden").CreateCollection(ctx, "user_actions", &options.CreateCollectionOptions{
		Capped:      &isCapped,
		SizeInBytes: &limitSize,
	})

	if err != nil {
		panic(err)
	}

	cappedCol := client.Database("garden").Collection("user_actions")

	for i := 0; i < 500; i++ {
		doc := bson.M{
			"username":    "kbanker",
			"action_code": rand.Intn(4),
			"time":        time.Now().UTC(),
			"n":           i,
		}

		cappedCol.InsertOne(ctx, doc)
	}
}
