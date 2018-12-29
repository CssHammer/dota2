package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/allbuleyu/dota2/config"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/options"
	"net/http"
	"strconv"
	"time"
)

func GetLastSeqNum() int64 {
	client, _ := config.NewMongoClient("")
	ctx, _ := context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	client.Connect(ctx)
	oneOptions := options.FindOne()
	oneOptions.Sort = bson.M{
		"_id":-1,
	}
	last_math_doc := client.Database("d2log").Collection("data_matches").FindOne(ctx, nil, oneOptions)

	match := new(ResultOfMatch)
	err := last_math_doc.Decode(match)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return match.Match_seq_num
}

func GetMatchHistoryBySeqNum(matches_requested int64)  {
	start_at_match_seq_num := GetLastSeqNum()

	addr := "http://api.steampowered.com/IDOTA2Match_570/GetMatchHistoryBySequenceNum/v1"
	querys := map[string]string{}
	querys["start_at_match_seq_num"] = strconv.FormatInt(start_at_match_seq_num, 10)
	querys["matches_requested"] = strconv.FormatInt(matches_requested, 10)
	addr = config.Addr(addr, querys)

	req, _ := http.NewRequest("GET", addr, nil)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http request err ", err)
		return
	}
	decoder := json.NewDecoder(resp.Body)

	res := struct {
		Result struct{
			Status int
			Matches []ResultOfMatch
		}
	}{}
	decoder.Decode(&res)

	if len(res.Result.Matches) == 0 {
		fmt.Println("get Matches = 0, so quit")
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	StoreMatches(ctx, res.Result.Matches)
}

func StoreMatches(ctx context.Context, matches []ResultOfMatch)  {
	// format data
	players := make([]PlayersOfMatch, 0)
	picksbans := make([]PicksBansOfMatch, 0)
	abilities := make([]AbilityUpgrades, 0)

	for i := range matches {
		match := &matches[i]
		match_id := match.Match_id
		match.Time_stamp = time.Now().Unix()

		for pi := range match.Players {
			player := &match.Players[pi]
			player.Match_id = match_id
			
			for ai := range player.Ability_upgrades {
				ability := player.Ability_upgrades[ai]
				ability.Match_id = match_id
				ability.Account_id = player.Account_id

				abilities = append(abilities, ability)
			}
			player.Ability_upgrades = nil
			
			players = append(players, *player)
		}

		for pbi := range match.Picks_bans {
			pickban := &match.Picks_bans[pbi]
			pickban.Match_id = match_id
			if pickban.Team == 0 {
				pickban.Team_id = match.Radiant_team_id
			}else if pickban.Team == 1 {
				pickban.Team_id = match.Dire_team_id
			}
			pickban.Leagueid = match.Leagueid

		}

		players = append(players, matches[i].Players...)
		picksbans = append(picksbans, matches[i].Picks_bans...)

		match.Players = nil
		match.Picks_bans = nil
	}


	client, err := config.NewMongoClient("")
	if err != nil {
		fmt.Println("client err", err)
		return
	}
	ctx,_ = context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("connect err", err)
		return
	}

	db := client.Database("d2log")

	// convert []type to []interface{}
	// convert start
	insertMatches := make([]interface{}, len(matches))
	insertPlayers := make([]interface{}, len(players))
	insertPicksBans := make([]interface{}, len(picksbans))
	insertAbilities := make([]interface{}, len(abilities))

	for i := range matches {
		insertMatches[i] = matches[i]
	}

	for i := range players {
		insertPlayers[i] = players[i]
	}

	for i := range picksbans {
		insertPicksBans[i] = picksbans[i]
	}

	for i := range abilities {
		insertAbilities[i] = abilities[i]
	}
	// convert start end

	resMatches, err := db.Collection("data_matches").InsertMany(ctx, insertMatches)
	if err != nil {
		fmt.Println("insert data_matches err", err)
		return
	}
	fmt.Println("isnerted data_matches:", len(resMatches.InsertedIDs))

	//ctxPlayers, _ := context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	//resPlayers, err := db.Collection("data_players").InsertMany(ctxPlayers, insertPlayers)
	resPlayers, err := db.Collection("data_players").InsertMany(ctx, insertPlayers)
	if err != nil {
		fmt.Println("insert data_players err", err)
		return
	}
	fmt.Println("isnerted data_players:", len(resPlayers.InsertedIDs))

	//ctxPicksBans, _ := context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	//resPicksBans, err := db.Collection("data_picksbans").InsertMany(ctxPicksBans, insertPicksBans)
	resPicksBans, err := db.Collection("data_picksbans").InsertMany(ctx, insertPicksBans)
	if err != nil {
		fmt.Println("insert data_picksbans err", err)
		return
	}
	fmt.Println("isnerted data_picksbans:", len(resPicksBans.InsertedIDs))

	//ctxAbilities, _ := context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	//resAbilities, err := db.Collection("data_ability_upgrades").InsertMany(ctxAbilities, insertAbilities)
	resAbilities, err := db.Collection("data_ability_upgrades").InsertMany(ctx, insertAbilities)
	if err != nil {
		fmt.Println("insert data_ability_upgrades err", err)
		return
	}
	fmt.Println("isnerted data_ability_upgrades:", len(resAbilities.InsertedIDs))
}
