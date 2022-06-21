package dealer

import (
	"github.com/againStore/data"
	"math/rand"
	"time"
)

type Dealer struct{}

func generateRandomPrice(down, up int) uint64 {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(up+down) - down
	return uint64(n)
}

func (d Dealer) Add(name string, quantity uint64) bool {
	a := generateRandomPrice(1, 20)
	b := a
	a += a * 2 / 10
	if b == a {
		a++
	}
	bul, err := data.Add(name, quantity, a, b)
	if err != nil {
		panic(err)
	}
	return bul
}
