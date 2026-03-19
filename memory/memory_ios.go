//go:build ios

package memory

import (
	"os"
	"runtime/debug"
	"runtime"

	// "time"
)

// const (
// 	// interval = 5
// 	// 33M
// 	maxMemory = 30 * 1024 * 1024
// )

// func forceFree(interval time.Duration) {
// 	go func() {
// 		for {
// 			time.Sleep(interval)
// 			debug.FreeOSMemory()
// 		}
// 	}()
// }

func InitForceFree() {
	os.Setenv("GODEBUG", "tlsmlkem=0")

	debug.SetGCPercent(30)
	// debug.SetMemoryLimit(maxMemory)
	// set golang threads num
	runtime.GOMAXPROCS(3)
	// duration := time.Duration(interval) * time.Second
	// forceFree(duration)
}
