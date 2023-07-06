package helper

import (
	"log"
	"os"
	"strings"
)

func GetWhiteListUsernames() map[string]bool {
	file, err := os.ReadFile(CONFIG.WhiteListTXT)
	if err != nil {
		log.Fatalf("Error reading whitelist file from config: %s", err.Error())
	}

	whiteListUsernames := make(map[string]bool)

	for _, username := range strings.Split(string(file), "\n") {
		whiteListUsernames[username] = true
	}

	return whiteListUsernames
}
