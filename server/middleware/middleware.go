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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var collection *mongo.Collection

func init (){
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv(){
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("error loading env file :(")
	}
}

func createDBInstance(){
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)
	client,err := mongo.Connect(context.TODO(), clientOptions)
	if err!=nil{
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("connected to mongoDB")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("collection instance created")
}

func GetAllCards (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := getAllCards()
	json.NewEncoder(w).Encode(result)
}

func GetOneCard (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := getOneCard()
	json.NewEncoder(w).Encode(result)
}

func CreateCard (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")	

	var card models.Card
	json.NewDecoder(r.Body).Decode(&card)
	createCard(card)
	json.NewEncoder(w).Encode(card)
}

func CardKnown (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")	
	
	params := mux.Vars(r)
	cardKnown(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoKnown (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoKnown(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteCard (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	deleteOneCard(params["id"])
}
func DeleteAllCards (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	deleteAll := deleteAllCards()
	json.NewEncoder(w).Encode(deleteAll)
}

func getAllCards() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}
	
	var results []primitive.M
	for cur.Next(context.Background()){
		var result bson.M
		e := cur.Decode(&result)
		if e!=nil{
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

func cardKnown(card string) {
	id, _ := primitive.ObjectIDFromHex(card)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"known": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

func createCard(card models.Card) {
	insertResult, err := collection.InsertOne(context.Background(), card)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single record", insertResult.InsertedID)
}

func undoKnown(card string) {
	id, _ := primitive.ObjectIDFromHex(card)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"known": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneCard(card string) {
	id, _ := primitive.ObjectIDFromHex(card)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("deleted item: ", d.DeletedCount)
}

func deleteAllCards() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("deleted", d.DeletedCount)
	return d.DeletedCount
}