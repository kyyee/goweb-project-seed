package xsecurity

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	result, err := GetJasypt(Options{}).Encrypt("123456")
	assert.Nil(t, err)
	log.Printf("encrypt password 123456, result is %s", result)
	assert.NotNil(t, result)
}

func TestDecrypt(t *testing.T) {
	result, err := GetJasypt(Options{}).Decrypt("M+yAiu2BnCFt0noSP/24xqg9V61T8zM/")
	assert.Nil(t, err)
	log.Printf("encrypt password 123456, result is %s", result)
	assert.NotNil(t, result)
	assert.Equal(t, result, "123456")
}
