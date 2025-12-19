package stealth

import (
	"math/rand"
	"time"
)

func HumanDelay() {
	time.Sleep(time.Duration(rand.Intn(800)+400) * time.Millisecond)
}
