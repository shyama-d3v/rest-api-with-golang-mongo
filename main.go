package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shyamapatro/restapi/router"
)

func main() {

	fmt.Println("MongoDB api")
	r := router.Router()
	fmt.Println("Server is getting started.......")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("listening at ported 3000.")

}
