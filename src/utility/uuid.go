package utility

import "regexp"

func IsUUID(uuid string) bool {
	noDashMatch, _ := regexp.MatchString(`^[0-9a-f]{32}$`, uuid)
	return noDashMatch
}
