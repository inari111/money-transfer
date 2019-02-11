package main

import (
	"log"
	"net/http"

	"github.com/inari111/money-transfer-study/di"
)

func main() {
	http.Handle("/", di.InitializeAPIHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
