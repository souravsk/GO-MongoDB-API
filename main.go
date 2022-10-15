package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/souravsk/GO-MongoDB-API/router"
)

func main() {
	fmt.Println("This is ManogDB API")
	r := router.Router() //calling the router
	fmt.Println("Server is getting to Started.....")
	log.Fatal(http.ListenAndServe(":4000", r)) //create the server at post 4000
	fmt.Println("Listening at port 4000 ...")
}
