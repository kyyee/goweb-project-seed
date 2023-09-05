package xerrors

import (
	"log"
	"sort"
	"testing"
)

func TestNewError(t *testing.T) {
	err := NewAssignCodeError(SYS_INTERNAL_ERROR)
	log.Printf("err: %+v", err)
	keys := make([]string, 0)
	for k, _ := range Msg {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	log.Printf("key size is: %d", len(keys))
	for _, k := range keys {
		log.Printf("key: %+v, value: %+v", ERROR_CODE_PRIFIX+k, Msg[k])
	}
}
