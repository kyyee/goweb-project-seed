package xsecurity

import (
	"log"
	"regexp"
)

func FilterEncryptString(originStr string) string {
	reg := regexp.MustCompile(`ENC\(([^\s]+)\)`)
	subStr := reg.FindStringSubmatch(originStr)
	if subStr != nil {
		raw, err := GetJasypt(Options{}).Decrypt(subStr[1])
		if err != nil {
			log.Fatalf("Decrypt password failed. %s", err)
			return ""
		}
		return reg.ReplaceAllString(originStr, raw)
	} else {
		return originStr
	}
}
