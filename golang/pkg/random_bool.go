package pkg

import (
	"math/rand"
	"time"
)

func RandomBool(scarcity int) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(scarcity) == 1
}