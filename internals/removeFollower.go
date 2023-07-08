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

		friendship, err := f.GetFriendship()
		if err != nil {
			log.Printf("Error getting friendship status for %s: %s", f.Username, err.Error())
			continue
		}

		if helper.CONFIG.ExcludeFollowing && friendship.Following {
			log.Printf("Skipping %s as you are following them", f.Username)
			continue
		}

		if p.WhiteListUsernames[f.Username] != true {
			log.Printf("Adding %s to list of followers to remove", f.Username)
			FollowerTORemove = append(FollowerTORemove, f)
			time.Sleep(1 * time.Second)
		} else {
			log.Printf("Skipping %s as they are in whitelist", f.Username)
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

		if rFollowers >= helper.CONFIG.Limit {
			log.Printf("Reached limit of %d removes", helper.CONFIG.Limit)
			break
		}

		params := &AccountNotWhiteListParams{
			WhiteListUsernames: p.WhiteListUsernames,
			Followers:          f.Users,
		}

		followers := GetNonWhiteListedAccount(params)

		if len(followers) != 0 {

			for _, follower := range followers {
				if follower != nil {

					err := follower.Block(false)
					err = follower.Unblock()
					if err != nil {
						log.Printf("Error remove follower %s: %s", follower.Username, err.Error())
					} else {
						log.Printf("Successfully remove followe %s", follower.Username) // Logging for successful unfollowing
					}
					rFollowers++
					sTime := rand.Intn(20) + 5
					log.Printf("Sleeping for %d seconds before removing next follower", sTime)
					time.Sleep(time.Duration(sTime) * time.Second)
				}
			}
		}
		next := f.Next()

		if !next {
			log.Printf("No more followers to unfollow")
			break
		}

		if rFollowers != 0 {
			//SleepTime is the time between each unfollowing
			randSTime := time.Duration(rand.Intn(50)+helper.CONFIG.Interval) * time.Second
			log.Printf("Sleeping for %s before starting next batch of unfollowing", randSTime)
			time.Sleep(randSTime)
		}
	}

	log.Printf("Unfollowed all non-whitelisted followers")
}
