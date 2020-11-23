package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/marcodeltorob/golang-bootcamp-2020/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func GetConnectionContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func NewDatabase(user, pass, host, dbname string) MongoDB {

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", user, pass, host, dbname)

	client, _ := mongo.Connect(GetConnectionContext(), options.Client().ApplyURI(uri))

	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())

	}

	log.Print("Connected to MongoDB!")

	return MongoDB{client}
}

func (mdb MongoDB) GetLocationById(woeid int) (*model.Location, error) {
	collection := mdb.client.Database("weather-app").Collection("locations")
	ctx := GetConnectionContext()

	location := &model.Location{}

	err := collection.FindOne(ctx, bson.M{"woeid": woeid}).Decode(location)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return location, nil

}

func (mdb MongoDB) GetAllLocations() ([]model.Location, error) {
	collection := mdb.client.Database("weather-app").Collection("locations")
	ctx := GetConnectionContext()
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(ctx)

	var locations []model.Location
	var location model.Location

	for cursor.Next(ctx) {
		cursor.Decode(&location)

		if cursor.Err() != nil {
			log.Println("Error al traer un documento: %w", cursor.Err())
		} else {
			locations = append(locations, location)
		}
	}

	return locations, nil
}
