package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/briannag31/mv-project3-quiz/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cardsCollection *mongo.Collection
var decksCollection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()

}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env file :(")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongoDB")
	database := client.Database("quiz-app-2")
	cardsCollection = database.Collection("cards")
	decksCollection = database.Collection("decks")
	fmt.Println("collection instance created")
}

func GetAllCards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payloadCards := getAllCards()
	json.NewEncoder(w).Encode(payloadCards)
}
func GetAllDecks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllDecks()
	json.NewEncoder(w).Encode(payload)
}
// func GetOneCard(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	params := mux.Vars(r)
// 	getOneCard(params["id"])
// 	json.NewEncoder(w).Encode(params["id"])
// 	// result := getOneCard()
// 	// json.NewEncoder(w).Encode(result)
// }

// func GetOneDeck(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	params := mux.Vars(r)
// 	getOneDeck(params["id"])
// 	json.NewEncoder(w).Encode(params["id"])
// }

func CreateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var card models.Card
	json.NewDecoder(r.Body).Decode(&card)
	createCard(card)
	json.NewEncoder(w).Encode(card)
}
func CreateDeck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var deck models.Deck
	json.NewDecoder(r.Body).Decode(&deck)
	createDeck(deck)
	json.NewEncoder(w).Encode(deck)

}

func CardKnown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	cardKnown(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoKnown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoKnown(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteOneCard(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
// func DeleteAllCards(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	deleteAll := deleteAllCards()
// 	json.NewEncoder(w).Encode(deleteAll)
// }

func getAllCards() []primitive.M {
	cursor, err := cardsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var cards []primitive.M
	for cursor.Next(context.Background()) {
		var card bson.M
		if err = cursor.Decode(&card); err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
		fmt.Println(card)
	}
	return cards
}

func getAllDecks() []primitive.M {
	cursor, err := decksCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var decks []primitive.M
	for cursor.Next(context.Background()) {
		var deck bson.M
		if err = cursor.Decode(&deck); err != nil {
			log.Fatal(err)
		}
		decks = append(decks, deck)
		fmt.Println(deck)
	}
	return decks
}

// func getOneCard(card string) {
// 	id, _ := primitive.ObjectIDFromHex(card)
	// var id primitive.ObjectID
	// filterCursor, err := cardsCollection.Find(context.Background(), bson.M{"_id": id})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var cardsFiltered []bson.M
	// if err = filterCursor.All(context.TODO(), &cardsFiltered); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(cardsFiltered)
	// opts := options.FindOne().SetSort( bson.M{"_id": id})
// 	var result bson.M
// 	filter := bson.D{{"_id", id}}
// 	err := cardsCollection.FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return
// 		}
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("found document %v", result)
// }
// func getOneDeck(deck string) {
// 	id, _ := primitive.ObjectIDFromHex(deck)
// 	filterCursor, err := decksCollection.Find(context.Background(), bson.M{"_id": id})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var decksFiltered []bson.M
// 	if err = filterCursor.All(context.Background(), &decksFiltered); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(decksFiltered)
// }

func cardKnown(card string) {
	id, _ := primitive.ObjectIDFromHex(card)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"known": true}}
	result, err := cardsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

func createCard(card models.Card) {
	insertResult, err := cardsCollection.InsertOne(context.Background(), card)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single record", insertResult.InsertedID)
}
func createDeck(deck models.Deck) {
	insertResult, err := decksCollection.InsertOne(context.Background(), deck)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single record", insertResult.InsertedID)
}

func undoKnown(card string) {
	id, _ := primitive.ObjectIDFromHex(card)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"known": false}}
	result, err := cardsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneCard(card string) {
	id, _ := primitive.ObjectIDFromHex(card)
	fmt.Println(id)
	fmt.Println(card)
	filter := bson.M{"_id": id}
	d, err := cardsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted item: ", d.DeletedCount)
}

// func deleteAllCards() int64 {
// 	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("deleted", d.DeletedCount)
// 	return d.DeletedCount
// }
