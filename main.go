package main

import (
	"fmt"
	"net/http"
	"log"
	"context"
	"time"
	"encoding/json"
	ioutil "io/ioutil"

	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type UserStructure struct {
	Email string
	Password string
}

func main () {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("ERROR:", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("ERROR:", err)
	} else {
		fmt.Println("connected to mongo")
	}
	// err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("ERROR:", err)
	}

	collection := client.Database("task-manager-api").Collection("users")
	fmt.Println(collection)

	var login UserStructure
	http.Handle("/login", &login)
	log.Fatal(http.ListenAndServe(":8080", nil))
	defer client.Disconnect(ctx)
}

func (login *UserStructure) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(data), &login)

	if err != nil {
		fmt.Println("ERROR:", err)
	}

	doSomething(login)

	w.Write(data)
}

func doSomething (login *UserStructure) {
	fmt.Println(*login)
}
