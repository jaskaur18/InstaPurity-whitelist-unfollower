package internal

import (
	"fmt"
	"github.com/jaskaur18/InstaPurity-whitelist-unfollower/helper"
	"log"
	"os"
	"path"

	"github.com/Davincible/goinsta"
	"github.com/jaskaur18/InstaPurity-whitelist-unfollower/types"
)

func CheckSessionExists(sessionPath string) bool {
	if _, ok := os.Stat(sessionPath); ok == nil {
		return true
	}
	return false
}

func LoginInsta(params *types.LoginParams) (insta *goinsta.Instagram, err error) {
	if params.Username == "" || params.Password == "" {
		return nil, types.ErrUsernamePassword
	}

	sessionPath := path.Join(helper.HomePath, "sessions", fmt.Sprintf(".%s", params.Username))
	_ = os.MkdirAll(path.Dir(sessionPath), os.ModePerm)

	if CheckSessionExists(sessionPath) {
		log.Printf("Session exists for %s", params.Username)
		insta, err = goinsta.Import(sessionPath)
		if err != nil {
			log.Printf("Error importing session for %s: %s", params.Username, err.Error())
			_ = os.Remove(sessionPath)
			insta = goinsta.New(params.Username, params.Password)
		}

		err := insta.OpenApp()
		if err != nil {
			return nil, types.LoginError{
				Msg: fmt.Sprintf("Error opening app: %s", err.Error()),
			}
		}

		return insta, nil
	} else {
		log.Printf("Session does not exist for %s", params.Username)
		insta = goinsta.New(params.Username, params.Password)
	}

	if insta == nil {
		log.Printf("Logging though password for %s", params.Username)

		if err := insta.Login(); err != nil {
			if err == goinsta.Err2FARequired {
				return nil, types.ErrLogin2FARequired
			}

			return nil, types.LoginError{
				Msg: err.Error(),
			}
		}
	}

	defer func(insta *goinsta.Instagram, path string) {
		err := insta.Export(path)
		if err != nil {
			log.Printf("Error exporting session for %s: %s :%s", params.Username, path, err.Error())
		}
	}(insta, sessionPath)

	return insta, nil
}
