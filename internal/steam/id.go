package steam

import (
	"regexp"
	"strconv"
)

var (
	steamIdPattern = regexp.MustCompile(`^STEAM_[0-5]:[01]:\d+$`)
)

func IsSteamID(provided string) bool {
	return steamIdPattern.MatchString(provided)
}

func IsSteamID64(provided string) bool {
	res, err := strconv.Atoi(provided)

	if err != nil {
		return false
	}

	return 0x0110000100000001 < res && res < 0x01100001FFFFFFFF
}
