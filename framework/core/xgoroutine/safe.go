package xgoroutine

import (
	"log"
	"runtime"
)

func Recover(fns ...func()) {
	for _, c := range fns {
		c()
	}
	if p := recover(); p != nil {
		log.Fatalf("recover a panic, %v", p)
	}
}

func RunSafe(fn func()) {
	defer Recover()
	fn()
}

func GoSafe(fn func()) {
	go RunSafe(fn)
}

func StackTrace() string {
	buf := make([]byte, 1<<11)
	i := runtime.Stack(buf, false)
	return string(buf[:i])

}
