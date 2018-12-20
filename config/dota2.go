package config

import (
	"net/url"
	"os"
)

const (
	JsonFormat string = "json"
	XMLFormat string = "xml"
)

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

// request url parse and format
func Addr(addr string, querys map[string]string) string {
	u, _ := url.Parse(addr)
	query := u.Query()
	query.Add("key", GetWebApiKey())
	query.Add("format", JsonFormat)

	for k, v := range querys {
		query.Add(k, v)
	}

	u.RawQuery = query.Encode()
	addr = u.String()

	return addr
}