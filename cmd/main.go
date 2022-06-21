package main

import (
	"fmt"
	"github.com/againStore/boss"
	"github.com/againStore/inventory"
	"github.com/againStore/product"
	"github.com/againStore/store"
	"strconv"
)

var List = product.List{}

func main() {
	boss := manager()
	start(boss)
}

func manager() boss.Boss {
	fmt.Print("What is your name?: ")
	sentence, er := inventory.RightInput()
	if er != nil {
		fmt.Errorf("error")
		return boss.Boss{}
	}
	sentence = sentence[:len(sentence)-1]
	name := sentence
	fmt.Print("Set password: ")
	sentence, er = inventory.RightInput()
	if er != nil {
		fmt.Errorf("error")
		return boss.Boss{}
	}
	sentence = sentence[:len(sentence)-1]
	password := sentence
	fmt.Print("Budget of your store: ")
	sentence, er = inventory.RightInput()
	if er != nil {
		fmt.Errorf("error")
		return boss.Boss{}
	}
	sentence = sentence[:len(sentence)-1]
	budget, er := strconv.Atoi(sentence)
	if er != nil {
		panic(er)
	}
	return boss.NewBoss(name, password, uint64(budget))
}

func start(b boss.Boss) {
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
			boss.MainBoss(b)
		} else if x == 2 {
			s := store.Store{}
			s.Sell()
		} else {
			return
		}
	}
}
