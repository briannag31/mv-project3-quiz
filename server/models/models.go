package models

import "go.mongodb.org/mongo-driver/bson/primitive"
// further updates: make collection its own struct and connect it to card
type Card struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Collection string             `json:"collection,omitempty"`
	Term       string             `json:"term,omitempty"`
	Answer     string             `json:"answer,omitempty"`
	Known      bool               `json:"known,omitempty"`
}
