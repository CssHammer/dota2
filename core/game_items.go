package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/allbuleyu/dota2/config"
	"net/http"
	"net/url"
)

type GameItems struct {
	Id int64
	// The item name by english
	OriginName string `json:"name"`

	// The in-game gold cost of the item.
	Cost int64

	// if the item is only available in the secret shop. 1 true 0 false
	Secret_shop uint8

	//  if the item is available in the side shop. 1 true 0 false
	Side_shop uint8

	// if the item is a recipe type item.
	Recipe int64

	// The localized name of the hero for use in name display. You will get it only if specifie 'language' parameter.
	Localized_name string
}

func GetGameItems() {

	// step1 parse addr and send request
	addr := "http://api.steampowered.com/IEconDOTA2_570/GetGameItems/v1"
	u, _ := url.Parse(addr)
	query := u.Query()
	query.Add("key", config.GetWebApiKey())
	query.Add("language", "zh")
	u.RawQuery = query.Encode()
	addr = u.String()

	req, _ := http.NewRequest("GET", addr, nil)
	req.Header.Add("name", "My name is mu!")

	httpDefaultClient := http.DefaultClient
	resp, err := httpDefaultClient.Do(req)
	if err != nil {
		fmt.Println("http request err: ", err)
		return
	}

	// step 2 parse json
	decoder := json.NewDecoder(resp.Body)
	jsonRes := struct {
		Result struct{
			Items []GameItems
			Status int64
		}
	}{}

	err = decoder.Decode(&jsonRes)
	if err != nil {
		fmt.Println("decode json err:", err)
		return
	}
	items := jsonRes.Result.Items

	// step 3 store or update items
	client, err := config.NewMongoClient("")
	if err != nil {
		fmt.Println("client err", err)
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("connect err", err)
		return
	}



	insertData := make([]interface{}, 0)
	// 更新,以后再说
	//updateData := make([]interface{}, 0)

	db := client.Database("d2log")

	//cur, err := db.Collection("data_items").Find(ctx, "")
	//db.Collection("data_items").UpdateMany()
	//for cur.Next(ctx) {
	//
	//}


	for i := range items {
		insertData = append(insertData, items[i])
	}

	resItem, err := db.Collection("data_items").InsertMany(ctx, insertData)
	if err != nil {
		fmt.Println("insert err", err)
		return
	}

	fmt.Println("insert rows = ",len(resItem.InsertedIDs))
}

