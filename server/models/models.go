package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// further updates: make collection its own struct and connect it to card

type Card struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Deck   primitive.ObjectID `json:"deck,omitempty" bson:"deck,omitempty"`
	Term   string             `json:"term,omitempty"`
	Answer string             `json:"answer,omitempty"`
	Known  bool               `json:"known,omitempty"`
}
type Deck struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DeckName string             `json:"deckName,omitempty"`
	Cards    Card
}

// Need to do:
// figure out how to call on deck and get all associated cards
