package types

import (
	"fmt"
	"github.com/Davincible/goinsta"
)

type Config struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Interval     int    `yaml:"interval"`
	Limit        int    `yaml:"limit"`
	WhiteListTXT string `yaml:"whitelist"`
}

type LoginParams struct {
	Username string
	Password string
}

type RemoveFollowerParams struct {
	*goinsta.Instagram
	WhiteListUsernames map[string]bool
}

func main() {

	var zList = []string{}
	var ylist = []string{}

	for _, x := range zList {
		for _, y := range ylist {
			if x == y {
				fmt.Println("Match found")
			}
		}
	}
}
