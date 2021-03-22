package main

import (
	"context"
	"fmt"
	"log"
	"loso/database"
)

var CNX = database.Connection()
var Conn, _ = database.NewCon("ln-smt")

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {

	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}

	colla := Conn.DB.Collection("test")
	insertResult1, err1 := colla.InsertOne(context.TODO(), ash)
	if err1 != nil {
		log.Fatal(err1)
	}

	insertResult2, err := colla.InsertOne(context.TODO(), misty)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document2: ", insertResult1.InsertedID)
	fmt.Println("Inserted a single document2: ", insertResult2.InsertedID)

}
