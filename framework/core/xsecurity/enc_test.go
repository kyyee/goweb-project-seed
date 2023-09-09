package xsecurity

import (
	"log"
	"testing"
)

func TestEnc(t *testing.T) {
	pwd := "ENC(VX04Iz9jYLcHEDF7a8sOBFBKPlJ3R+nz)"
	decode := FilterEncryptString(pwd)
	log.Printf("decode password is %s", decode)
}
