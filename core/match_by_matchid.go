package core

import (
	"encoding/json"
	"fmt"
	"github.com/allbuleyu/dota2/config"
	"github.com/allbuleyu/dota2/enum"
	"net/http"
	"net/url"
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
	Duration int

	//
	Pre_game_duration int

	// Unix timestamp of when the match began.
	Start_time int

	// The matches unique ID.
	Match_id int

	// A 'sequence number', representing the order in which matches were recorded.
	Match_seq_num int

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

	// A list of the picks and bans in the match, if the game mode is Captains Mode. (Ranked Matchmaking also has)
	Picks_bans []PicksBansOfMatch

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
}

func GetMatchDetail(matchid string) {
	// step 1	?key=D524A0B32AECE6B5986B5EFCF09AA58D&match_id=4267110473
	addr := "http://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/v1"

	// add query
	u, err := url.Parse(addr)
	query := u.Query()
	query.Add("key", config.GetWebApiKey())
	query.Add("match_id", matchid)
	u.RawQuery=query.Encode()
	addr = u.String()

	// step 2  format the parameter
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		panic(err)
	}
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()


	//result := map[string]interface{}{
	//	"result":ResultOfMatch{},
	//}

	res := struct {
		Result ResultOfMatch
	}{}

	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		fmt.Println("decode false by:", err)
		panic(err)
	}

	// mongo
	client, err := config.NewMongoClient("")
	client.Database("")

	// not done
	// store match detail
	match := res.Result


	// store players detail
	players := res.Result.Players

	// https://yq.aliyun.com/articles/69662 没弄明白最后一个例子


}