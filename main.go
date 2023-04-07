package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kscastro/todo-list-go/src/router"
)

func main() {
	r := router.Router()

	fmt.Println("Starting port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
