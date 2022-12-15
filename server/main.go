package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/briannag31/mv-project3-quiz/router"
)


func main() {
	r := router.Router()
	fmt.Println("starting the server on port 3001")
	log.Fatal(http.ListenAndServe(":3001", r))

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// connectionString := os.Getenv("DB_URI")
	// clientOptions := options.Client().ApplyURI(connectionString)

	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Disconnect(ctx)

	// database := client.Database("quiz-app-2")
	// cardsCollection := database.Collection("cards")
	// categoriesCollection := database.Collection("categories")

	// categoryResult, err := categoriesCollection.InsertOne(ctx, bson.D{
	// 	{"group", "Group 1"},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(categoryResult.InsertedID)

	// cardResult, err := cardsCollection.InsertMany(ctx, []interface{}{
	// bson.D{
	// 	{"category", categoryResult.InsertedID},
	// 	{"term", "test3"},
	// 	{"answer", "this is a test three"},
	// 	{"known", false},
	// },
	// bson.D{
	// 	{"category", categoryResult.InsertedID},
	// 	{"term", "test4"},
	// 	{"answer", "this is a test four"},
	// 	{"known", false},
	// },
	// bson.D{
	// 	{"category", categoryResult.InsertedID},
	// 	{"term", "test5"},
	// 	{"answer", "this is a test five"},
	// 	{"known", false},
	// },	
	// } )
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(cardResult.InsertedIDs)

}
