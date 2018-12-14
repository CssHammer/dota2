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
	if err != nil {
		return "", err
	}
	query := u.Query()

	query.Add("key", config.GetWebApiKey())
	query.Add("format", "json")

	return "", nil
}