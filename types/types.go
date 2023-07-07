package types

import (
	"github.com/Davincible/goinsta"
)

type Config struct {
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	Interval         int    `yaml:"interval"`
	Limit            int    `yaml:"limit"`
	ExcludeFollowing bool   `yaml:"exclude_following"`
	WhiteList        string `yaml:"whitelist"`
}

type LoginParams struct {
	Username string
	Password string
}

type RemoveFollowerParams struct {
	*goinsta.Instagram
	WhiteListUsernames map[string]bool
}
