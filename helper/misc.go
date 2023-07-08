package helper

import (
	"log"
	"os"
	"path"
	"strings"
)

func GetWhiteListUsernames() map[string]bool {
	whiteListFilePath := path.Join(HomePath, CONFIG.WhiteList)

	file, err := os.ReadFile(whiteListFilePath)
	if err != nil {
		log.Fatalf("Error reading whitelist file from config: %s", err.Error())
	}

	whiteListUsernames := make(map[string]bool)

	for _, username := range strings.Split(string(file), "\n") {
		username = strings.TrimSpace(username)
		username = strings.ReplaceAll(username, "@", "")

		whiteListUsernames[username] = true
	}

	return whiteListUsernames
}
