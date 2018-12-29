package main

import (
	"github.com/allbuleyu/dota2/core"
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

	core.GetMatchHistoryBySeqNum(100)
	return
	//core.GetMatchDetail(4262769848)
	//core.GetGameItems()

	core.GetTeamsInfo()
}



