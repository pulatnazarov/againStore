package main

import (
	"fmt"
	"github.com/againStore/inventory"
	"github.com/againStore/product"
	"github.com/againStore/store"
	"strconv"
)

var List = product.List{}

func main() {
	start()
}

func start() {
	for {
	end:
		fmt.Print(`
		1 - Boss
		2 - Customer
		3 - Exit
	`)
		sentence, er := inventory.RightInput()
		if er != nil {
			fmt.Errorf("error: %d", er)
			goto end
		}
		sentence = sentence[:len(sentence)-1]
		x, er := strconv.Atoi(sentence)
		if er != nil {
			fmt.Printf("error: %d", er)
			goto end
		}
		if x > 3 || x < 1 {
			fmt.Println("Enter one of them")
			goto end
		}
		if x == 1 {
			//boss.Bossmain()
		} else if x == 2 {
			s := store.Store{}
			s.Sell()
		} else {
			return
		}
	}
}
