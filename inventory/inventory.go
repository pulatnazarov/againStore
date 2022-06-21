package inventory

import (
	"bufio"
	"fmt"
	"github.com/againStore/data"
	"os"
	"strconv"
)

type Inventory struct {
}

func RightInput() (string, error) {
	sentence, er := bufio.NewReader(os.Stdin).ReadString('\n')
	return sentence, er
}

func (i Inventory) AddOrIncrease() {
again:
	fmt.Print("Name of product: ")
	sentence, er := RightInput()
	if er != nil {
		fmt.Errorf("error:", er)
		goto again
	}
	name := sentence[:len(sentence)-1]
quan:
	fmt.Print("Quantity of product: ")
	sentence, err := RightInput()
	if err != nil {
		fmt.Errorf("error", er)
		goto quan
	}
	sentence = sentence[:len(sentence)-1]
	quantity, er := strconv.Atoi(sentence)
	if er != nil {
		fmt.Errorf("Error", er)
	}
	if quantity <= 0 {
		fmt.Println("Enter positive number!")
		goto quan
	}
	dealer.Add(name, quantity)
}

func Sell(name string, quantity uint64) bool {
	price, b := data.Sell(name, quantity)
	if b {
		basket.Cash(name, quantity, price)
		fmt.Println("Added to your basket!")
		return true
	}
	fmt.Println("We dont have or there is not enough amount of this product!\nWe will bring it next time!")
	dealer.Add(name, quantity)
	return false
}

func (i Inventory) Delete() {
	
}
