package feeder

import "regexp"

func validFormat(prod string) bool {
	r, _ := regexp.Compile(`([A-Z]{4})-(\d{4})`)
	return r.MatchString(prod)
}
