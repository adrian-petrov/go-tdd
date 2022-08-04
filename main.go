package main

import (
	"gotdd/dependencyinjection"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreeterHandler)))
}
