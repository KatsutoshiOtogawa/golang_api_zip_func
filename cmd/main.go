package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	// Blank-import the function package so the init() runs
	_ "source.developers.google.com/p/forlesson/r/golang_api_function.git"
)

func main() {

	port := os.Getenv("PORT")

	// if you don't set PORT variable, error and exit.
	if port == "" {
		log.Fatal("requre set environment variable PORT.")
	}

	err := funcframework.Start(port)
	if err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
