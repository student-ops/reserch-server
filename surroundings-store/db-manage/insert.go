package dbmanage

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Surroundings struct {
	DeviceId    int
	Tempreture  float32
	Airpressure float32
	Date        time.Time
}

func DbConnection() (err error, client *mongo.Client) {
	var cred options.Credential
	cred.AuthSource = "admin"
	cred.Username = "root"
	cred.Password = "password"
	// Rest of the code will go here
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred)

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err, client
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return err, client
	}

	fmt.Println("Connected to MongoDB!")
	return err, client
}

func InsertSurroundings(client *mongo.Client, data Surroundings) {
	collection := client.Database("test").Collection("surroundings")

	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
