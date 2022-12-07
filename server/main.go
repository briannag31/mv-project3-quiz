package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/briannag31/mv-project3-quiz/router"
)

func main (){
	r := router.Router()
	fmt.Println("starting the server on port 3001")
	log.Fatal(http.ListenAndServe(":3001", r))
}