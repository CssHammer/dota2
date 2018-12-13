package config

import "os"

//WEBAPIKEY
func GetWebApiKey() string {
	return os.Getenv("WEBAPIKEY")
}

//STEAMID
func GetSteamId() string {
	return os.Getenv("STEAMID")
}

//DOTA2Id
func GetD2Id() string {
	return os.Getenv("DOTA2Id")
}