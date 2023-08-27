package utils

import "time"

func Duration(s string) time.Duration {
	d, e := time.ParseDuration(s)
	if e != nil {
		// todo
	}
	return d
}
