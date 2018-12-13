package core

import (
	"github.com/allbuleyu/dota2/config"
	"net/url"
)

type CommonParam struct {
	Key string
	Format string

}

func commonFormatJsonParams(rawurl string) (string, error) {
	u, err := url.Parse(rawurl)
	query := u.Query()

	query.Add("key", config.GetWebApiKey())
	query.Add("format", "json")
}