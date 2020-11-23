package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	LocationType string             `json:"location_type" bson:"location_type"`
	Woeid        int                `json:"woeid" bson:"woeid"`
	LattLong     string             `json:"latt_long" bson:"latt_long"`
}
