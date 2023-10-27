package approxtime

import (
	"sync"
	"time"
)

var (
	now     time.Time
	nowLock sync.RWMutex
)

func Update() {
	nowLock.Lock()
	defer nowLock.Unlock()
	now = time.Now()
}

func run() {
	for {
		Update()
		time.Sleep(time.Millisecond * 10)
	}
}

func Now() time.Time {
	nowLock.RLock()
	defer nowLock.RUnlock()
	return now
}

func Since(t time.Time) time.Duration {
	n := Now()
	return n.Sub(t)
}

func init() {
	Update()
	go run()
}

