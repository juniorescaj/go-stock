package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/juniorescaj/go-stock/api/health"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start() {
	router := mux.NewRouter()

	client, err := mongodbConnect()

	if err != nil {
		panic(err)
	}

	health.Map(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func mongodbConnect() (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))

	if err != nil {
		return nil, err
	}

	return client, nil
}
