package store

import (
	"fmt"
	"github.com/againStore/basket"
	"github.com/againStore/boss"
	"strconv"

	"github.com/againStore/inventory"
)

type Store struct {
}

func (s Store) Sell() {
	var Bs basket.Trash
	for {
	end:
		fmt.Printf(`
		1-Shop
		2-List
		3-Exit
	`)
		sentence, er := inventory.RightInput()
		if er != nil {
			fmt.Println(" ", er)
			goto end
		}

		sentence = sentence[:len(sentence)-1]
		x, er := strconv.Atoi(sentence)
		if er != nil {
			fmt.Println(er)
			goto end
		}
		if x > 3 || x < 1 {
			fmt.Println(er)
			goto end
		}
		if x == 1 {
			s.Shop(&Bs)
		} else if x == 2 {
			Bs.Show()
		} else {
			Bs.Check()
			return
		}
	}
}

func (s Store) Shop(Bs *basket.Trash) {
	fmt.Print("Enter product name: ")
	sentence, er := inventory.RightInput()
	if er != nil {
		fmt.Println(er)
	}
	sentence = sentence[:len(sentence)-1]
start:
	fmt.Print("Enter product quantity(only number): ")
	quantity, er := inventory.RightInput()
	if er != nil {
		fmt.Println(er)
	}
	quantity = quantity[:len(quantity)-1]
	x, er := strconv.Atoi(quantity)
	if er != nil {
		fmt.Println(er)
		goto start
	}
	if x <= 0 {
		fmt.Println("enter positive number")
		goto start
	}
	price, originPrice, bul := inventory.Sell(sentence, uint64(x))
	if bul {
		boss.BassFile.Budget += originPrice * uint64(x)
		boss.BassFile.Profit += (price - originPrice) * uint64(x)
		Bs.Append(sentence, uint64(x), price)
	}
}
