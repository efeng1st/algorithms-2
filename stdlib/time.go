package stdlib

import (
	"log"
	"time"
)

// TimeTrack 检查函数运行时间
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
