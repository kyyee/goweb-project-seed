package utils

import (
	"fmt"
	"time"
)

func Duration(s string) time.Duration {
	d, e := time.ParseDuration(s)
	if e != nil {
		// todo
	}
	return d
}

func TimeToMicroSecondStr(cost time.Duration) string {
	return fmt.Sprintf("[%vms]", float64(cost.Microseconds()))
}

func TimeToSecondStr(cost time.Duration) string {
	return fmt.Sprintf("[%vs]", float64(cost.Microseconds())/1000)
}
