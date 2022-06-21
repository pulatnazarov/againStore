package store

import (
	"fmt"
	"strconv"

	"github.com/againStore/inventory"
)

type Store struct {
}

func (s Store) Sell() {
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
		if x > 2 || x < 1 {
			fmt.Println(er)
			goto end
		}
		if x == 1 {
			s.Shop()
		} else if x == 2 {

		}
	}
}
func (s Store) Shop() {
	fmt.Println("Mahsulot nomini kiriting: ")
	sentence, er := inventory.RightInput()
	if er != nil {
		fmt.Println(er)
	}
	sentence = sentence[:len(sentence)-1]
start:
	fmt.Println("Enter product qountity(only number): ")
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
	
}

func (s Store) List() {

}
