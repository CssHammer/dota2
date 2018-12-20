package main

import (
	"context"
	"fmt"
	"github.com/allbuleyu/dota2/core"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type S struct {
	Response R
}

type R struct {
	Game_count int64
	Games []Games
}

type Games struct {
	Appid int64
	Playtime_forever int64
}

func main() {
	//core.GetMatchDetail(4262769848)
	core.GetGameItems()
}


type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Raw map[string]interface{} `json:"raw"` // Raw data
}

type heros struct {
	Id int64
	Local_name string
	Name string
}

func findHeros() {
	var hero heros
	hero=heros{2, "xxx", "xxx"}
	client, err := mongo.NewClient("mongodb://192.168.1.90:27017")
	err = client.Connect(context.TODO())

	collection := client.Database("d2log").Collection("heros")
	//collection := mongo.Client.Database("d2log").Collection("bar")
	ctx := context.Background()
	result, err := collection.InsertOne(ctx, hero)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.InsertedID)
}