package main

import (
	"github.com/jaskaur18/InstaPurity-whitelist-unfollower/helper"
	internal "github.com/jaskaur18/InstaPurity-whitelist-unfollower/internals"
	"github.com/jaskaur18/InstaPurity-whitelist-unfollower/types"
	"log"
)

func StartUnfollowing() {
	whiteListUsernames := helper.GetWhiteListUsernames()

	log.Printf("WhiteListed Usernames: %v", whiteListUsernames)

	log.Printf("Starting Unfollowing...")

	p := &types.LoginParams{
		Username: helper.CONFIG.Username,
		Password: helper.CONFIG.Password,
	}
	g, err := internal.LoginInsta(p)
	if err != nil {
		log.Fatalf("Error logging in: %s", err.Error())
	}

	log.Printf("Logged in as %s", g.Account.Username)

	removeFollowerParams := &types.RemoveFollowerParams{
		Instagram:          g,
		WhiteListUsernames: whiteListUsernames,
	}

	internal.RemoveFollowers(removeFollowerParams)

}

func init() {
	_, err := helper.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err.Error())
	}
}

func main() {
	//menu := gocliselect.NewMenu("What You Want To Do?")
	//
	//menu.AddItem("Start Unfollowing", "unfollow")
	//menu.AddItem("exit", "exit")
	//
	//choice := menu.Display()
	//
	//switch choice {
	//case "unfollow":
	StartUnfollowing()
	//case "exit":
	//	log.Printf("Exiting...")
	//	return
	//}
}
