package main

import (
	"fmt"
	"github.com/allbuleyu/dota2/config"
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
	i := 9999

	and := func(x int) int {
		countX := 0
		for x>0 {
			countX++
			x=x&(x-1)
		}
		return countX
	}


	fmt.Println(fmt.Sprintf("%b", i), and(i))
	return

	log := config.Logger()

	log.Error("奥德赛发生的发送大发是打发斯蒂芬sss")
	return

	core.GetMatchHistoryBySeqNum(100)
	return
	//core.GetMatchDetail(4262769848)
	//core.GetGameItems()

	core.GetTeamsInfo()
}



