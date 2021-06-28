package main

import (
	"fmt"
	"learn-go/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server in 8080")

	log.Fatalf(http.ListenAndServe(":8080", r))
}
