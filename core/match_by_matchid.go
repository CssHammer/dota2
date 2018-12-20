package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/allbuleyu/dota2/config"
	"github.com/allbuleyu/dota2/enum"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"net/url"
	"strconv"
	"time"
)


//Tower Status
//A particular teams tower status is given as a 16-bit unsigned integer.
// The rightmost 11 bits represent individual towers belonging to that team;
// see below for a visual representation.
//
//┌─┬─┬─┬─┬─────────────────────── Not used.
//│ │ │ │ │ ┌───────────────────── Ancient Bottom
//│ │ │ │ │ │ ┌─────────────────── Ancient Top
//│ │ │ │ │ │ │ ┌───────────────── Bottom Tier 3
//│ │ │ │ │ │ │ │ ┌─────────────── Bottom Tier 2
//│ │ │ │ │ │ │ │ │ ┌───────────── Bottom Tier 1
//│ │ │ │ │ │ │ │ │ │ ┌─────────── Middle Tier 3
//│ │ │ │ │ │ │ │ │ │ │ ┌───────── Middle Tier 2
//│ │ │ │ │ │ │ │ │ │ │ │ ┌─────── Middle Tier 1
//│ │ │ │ │ │ │ │ │ │ │ │ │ ┌───── Top Tier 3
//│ │ │ │ │ │ │ │ │ │ │ │ │ │ ┌─── Top Tier 2
//│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ ┌─ Top Tier 1
//│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
//0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

//Barracks Status
//A particular teams tower status is given as an 8-bit unsigned integer.
// The rightmost 6 bits represent the barracks belonging to that team;
// see below for a visual representation.
//
//┌─┬───────────── Not used.
//│ │ ┌─────────── Bottom Ranged
//│ │ │ ┌───────── Bottom Melee
//│ │ │ │ ┌─────── Middle Ranged
//│ │ │ │ │ ┌───── Middle Melee
//│ │ │ │ │ │ ┌─── Top Ranged
//│ │ │ │ │ │ │ ┌─ Top Melee
//│ │ │ │ │ │ │ │
//0 0 0 0 0 0 0 0
type ResultOfMatch struct{
	// Dictates the winner of the match, true for radiant; false for dire.
	Radiant_win bool `json:"radiant_win"`

	//The length of the match, in seconds since the match began.
	Duration int64

	//
	Pre_game_duration int64

	// Unix timestamp of when the match began.
	Start_time int64

	// The matches unique ID.
	Match_id int64

	// A 'sequence number', representing the order in which matches were recorded.
	Match_seq_num int64

	// See #Tower Status below.
	Tower_status_radiant uint16
	Tower_status_dire uint16


	Barracks_status_radiant uint8
	Barracks_status_dire uint8

	// The server cluster the match was played upon. Used for downloading replays of matches.
	Cluster int

	// The time in seconds since the match began when first-blood occurred.
	First_blood_time int

	// See the enum package LobbyType
	Lobby_type enum.LobbyType

	// The amount of human players within the match.
	Human_players int

	// The league that this match was a part of. A list of league IDs can be found via the GetLeagueListing method.
	// https://wiki.teamfortress.com/wiki/WebAPI/GetLeagueListing
	Leagueid int

	// The number of thumbs-up the game has received by users.
	Positive_votes int

	// The number of thumbs-down the game has received by users.
	Negative_votes int

	// game_mode
	Game_mode enum.D2GameMode

	//
	Flags int

	//	0 - Source 1, 1 - Source 2
	Engine int

	// The team radiant was killed
	Radiant_score int

	// The team dire was killed
	Dire_score int

	//
	Radiant_team_id int

	Radiant_name string

	Radiant_logo int64

	Radiant_team_complete int

	// captain of radiant steamid
	Radiant_captain int64

	//
	Dire_team_id int

	Dire_name string

	Dire_logo int64

	Dire_team_complete int

	// captain of dire steamid
	Dire_captain int64

	// A list of the picks and bans in the match, if the game mode is Captains Mode. (Ranked Matchmaking also has)
	Picks_bans []PicksBansOfMatch `json:"picks_bans"`

	Players []PlayersOfMatch `json:"players"`
}

type PicksBansOfMatch struct {
	//Whether this entry is a pick (true) or a ban (false).
	Is_pick bool

	// The hero's unique ID. A list of hero IDs can be found via the GetHeroes method.
	Hero_id uint8
	// The team who chose the pick or ban; 0 for Radiant, 1 for Dire.
	Team int

	// The order of which the picks and bans were selected;0-21
	Order int

	// Match_id
	Match_id int64 `json:"-"`
}

//Player Slot
//A player's slot is returned via an 8-bit unsigned integer. The first bit represent the player's team,
// false if Radiant and true if dire.
// The final three bits represent the player's position in that team, from 0-4.
//┌─────────────── Team (false if Radiant, true if Dire).
//│ ┌─┬─┬─┬─────── Not used.
//│ │ │ │ │ ┌─┬─┬─ The position of a player within their team (0-4).
//│ │ │ │ │ │ │ │
//0 0 0 0 0 0 0 0
type PlayersOfMatch struct{
	// matchid
	Match_id int64 `json:"-"`

	Account_id int64
	Player_slot uint8
	Hero_id uint8
	Item_0 int
	Item_1 int
	Item_2 int
	Item_3 int
	Item_4 int
	Item_5 int
	Backpack_0 int
	Backpack_1 int
	Backpack_2 int
	Kills int
	Deaths int
	Assists int
	Leaver_status int
	Last_hits int
	Denies int
	Gold_per_min int
	Xp_per_min int
	Level int
	Hero_damage int
	Tower_damage int
	Hero_healing int
	Gold int
	Gold_spent int
	Scaled_hero_damage int
	Scaled_tower_damage int
	Scaled_hero_healing int

	//
	Ability_upgrades []AbilityUpgrades
}

type AbilityUpgrades struct {
	// Account_id
	Account_id int64 `json:"-"`

	// matchid
	Match_id int64	`json:"-"`

	// ability id
	Ability int64

	// upgrade time of ability
	Time time.Duration

	// level of hero
	Level int64
}

func GetMatchDetail(matchid int64) {
	// mongo
	client, err := config.NewMongoClient("")
	if err != nil {
		fmt.Println("new client is failure", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeOutDuration)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("client connect is failure", err)
		return
	}
	db := client.Database("d2log")
	collection := db.Collection("data_matchs")

	defer cancel()
	filter := bson.M{"Match_id": matchid}


	isExist := ResultOfMatch{}
	err = collection.FindOne(ctx, filter).Decode(&isExist)
	if err != nil && err != mongo.ErrNoDocuments {
		fmt.Println("find match is failure:", err)
		return
	}

	//if isExist.Match_id != 0 {		// exist
	//	fmt.Println("this match is exist, matchid = ", matchid)
	//	return
	//}

	// step 1	?key=D524A0B32AECE6B5986B5EFCF09AA58D&match_id=4267110473
	addr := "http://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/v1"

	// add query
	u, err := url.Parse(addr)
	query := u.Query()
	query.Add("key", config.GetWebApiKey())
	query.Add("match_id", strconv.FormatInt(matchid, 10))
	u.RawQuery=query.Encode()
	addr = u.String()

	// step 2  format the parameter
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		fmt.Println("http new request is failure:", err)
		return
	}
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("http get is failure:", err)
		return
	}

	// json decoder
	decoder := json.NewDecoder(resp.Body)

	// step 3
	// store match detail start
	res := struct {
		Result ResultOfMatch
	}{}
	err = decoder.Decode(&res)
	if err != nil {
		fmt.Println("decode failure:", err)
		return
	}

	match := res.Result

	players := match.Players
	picksbans := match.Picks_bans

	match.Players = nil
	match.Picks_bans = nil


	oneRes, err := collection.InsertOne(ctx, match)
	if err != nil {
		fmt.Println("match insert is failure:", err)
		return
	}
	fmt.Println(oneRes.InsertedID)
	// match end


	// store players start
	mData := make([]interface{}, 0)

	mAbilityUpData := make([]interface{}, 0)
	for i := range players {
		players[i].Match_id = matchid

		// 每个玩家所持英雄技能加点信息
		for _, v := range players[i].Ability_upgrades {
			v.Match_id = matchid
			v.Account_id = players[i].Account_id

			mAbilityUpData = append(mAbilityUpData, v)
		}

		players[i].Ability_upgrades = nil
		mData = append(mData, players[i])
	}

	rp, err := db.Collection("data_players").InsertMany(ctx, mData)
	if err != nil {
		fmt.Println("players insert failure: ", err)
		return
	}
	fmt.Println("_id:", rp.InsertedIDs)


	resAbility, err := db.Collection("data_ability_upgrades").InsertMany(ctx, mAbilityUpData)
	if err != nil {
		fmt.Println("data_ability_upgrades insert failure: ", err)
		return
	}
	fmt.Println("_id:", resAbility.InsertedIDs)

	// store players end

	// store picksbans start
	mData = make([]interface{}, 0)
	for i := range picksbans {

		picksbans[i].Match_id = matchid
		mData = append(mData, picksbans[i])
	}

	// the match not have picks bans (random draft)
	if len(picksbans) == 0 {
		return
	}

	respb, err := db.Collection("data_picksbans").InsertMany(ctx, mData)
	fmt.Println("_id:", respb.InsertedIDs)
	if err != nil {
		fmt.Println("picksbans insert failure: ", err)
		return
	}
	// store picksbans end

	// store

	// store end

}