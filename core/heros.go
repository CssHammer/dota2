package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/allbuleyu/dota2/config"
	"net/http"
)

type Heroes struct {
	Name string
	Id int64
	En_name string
	Zh_name string
	Localized_name string
}

func GetHeroes(language string) {
	addr := "https://api.steampowered.com/IEconDOTA2_570/GetHeroes/v1/"
	querys := map[string]string{
		"language": language,
	}
	addr = config.Addr(addr, querys)

	resp, err := http.Get(addr)
	if err != nil {
		fmt.Println("http query err", err)
		return
	}

	decoder := json.NewDecoder(resp.Body)
	result := struct {
		Result struct{
			Heroes []Heroes
			Status int64
			count int64
		}
	}{}

	err = decoder.Decode(&result)
	if err != nil {
		fmt.Println("decode err", err)
		return
	}

	heroes := result.Result.Heroes

	insertData := make([]interface{}, 0)
	for i := range heroes {
		if language == "zh" {
			heroes[i].Zh_name = heroes[i].Localized_name
		}else if language == "en" {
			heroes[i].En_name = heroes[i].Localized_name
		}


		insertData = append(insertData, heroes[i])
	}

	client, err := config.NewMongoClient("")
	ctx, _:= context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	err = client.Connect(ctx)

	res, err := client.Database("d2log").Collection("data_heroes").InsertMany(ctx, insertData)
	if err != nil {
		fmt.Println("insert err", err)
		return
	}
	fmt.Println(len(res.InsertedIDs))
}
