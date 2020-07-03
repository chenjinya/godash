package godash

import (
	"log"
	"math/rand"
	"time"
)

func RandSleep() {

	r := rand.New(rand.NewSource(time.Now().Unix()))
	seed := r.Int() % 1000
	log.Println("Rand Sleep:", seed)
	time.Sleep(time.Duration(seed) * time.Millisecond)
}
