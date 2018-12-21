package main

import (
	"context"
	"fmt"
	"github.com/allbuleyu/dota2/config"
	"github.com/allbuleyu/dota2/core"
	"github.com/mongodb/mongo-go-driver/bson"
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
	client, err := config.NewMongoClient("")
	if err != nil {

	}
	ctx, _ := context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	client.Connect(ctx)
	filter := bson.M{"team_id": map[string]interface{}{"$in":[]int64{2,3, 4}}}
	cur, err := client.Database("d2log").Collection("data_teams").Find(ctx, filter)
	defer cur.Close(ctx)
	teams := make([]core.TeamInfo, 0)

	for cur.Next(ctx) {
		var res core.TeamInfo

		err = cur.Decode(&res)
		teams = append(teams, res)
	}
	fmt.Println(len(teams), teams)

	updateData := make([]interface{}, 0)
	for i := range teams {
		teams[i].Name = teams[i].Name + "xxx"

		updateData = append(updateData, teams[i])
	}
	client.Database("d2log").Collection("data_teams").UpdateMany(ctx, "", updateData[0])
	return

	//core.GetMatchDetail(4262769848)
	//core.GetGameItems()

	core.GetTeamsInfo()
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