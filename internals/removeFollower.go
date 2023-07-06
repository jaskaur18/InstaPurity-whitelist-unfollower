package internal

import (
	"github.com/Davincible/goinsta"
	"github.com/jaskaur18/InstaPurity-whitelist-unfollower/helper"
	"github.com/jaskaur18/InstaPurity-whitelist-unfollower/types"
	"log"
	"math/rand"
	"time"
)

type AccountNotWhiteListParams struct {
	WhiteListUsernames map[string]bool
	Followers          []*goinsta.User
}

func GetNonWhiteListedAccount(p *AccountNotWhiteListParams) []*goinsta.User {
	var FollowerTORemove []*goinsta.User
	for _, f := range p.Followers {
		if p.WhiteListUsernames[f.Username] != true {
			FollowerTORemove = append(FollowerTORemove, f)
		}
	}
	return FollowerTORemove
}

func RemoveFollowers(p *types.RemoveFollowerParams) {

	log.Printf("Starting Unfollowing...")
	log.Printf("Getting followers...")

	f := p.Account.Followers("")
	rFollowers := 0
	for {

		params := &AccountNotWhiteListParams{
			WhiteListUsernames: p.WhiteListUsernames,
			Followers:          f.Users,
		}

		followers := GetNonWhiteListedAccount(params)

		if len(followers) != 0 {

			for _, follower := range followers {
				if follower != nil {

					err := follower.Unfollow()
					if err != nil {
						log.Printf("Error unfollowing %s: %s", follower.Username, err.Error())
					} else {
						log.Printf("Successfully unfollowed %s", follower.Username) // Logging for successful unfollowing
					}
					rFollowers++
					sTime := rand.Intn(10)
					log.Printf("Sleeping for %d seconds before unfollowing next user", sTime)
					time.Sleep(time.Duration(sTime) * time.Second)
				}
			}
		}
		next := f.Next()
		if !next {
			log.Printf("No more followers to unfollow")
			break
		}

		//SleepTime is the time between each unfollowing
		randSTime := time.Duration(rand.Intn(30)+helper.CONFIG.Interval) * time.Second
		log.Printf("Sleeping for %s before starting next batch of unfollowing", randSTime)
		time.Sleep(randSTime)
	}

	log.Printf("Unfollowed all non-whitelisted followers")
}
