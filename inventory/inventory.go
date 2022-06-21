package inventory

import (
	"bufio"
	"fmt"
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
new:
	fmt.Print(`
	1 - Add
	2 - Increase
`)
	sentence, er := RightInput()
	if er != nil {
		fmt.Errorf("err", er)
		goto new
	}
	sentence = sentence[:len(sentence)-1]
	x, er := strconv.Atoi(sentence)
	if er != nil {
		fmt.Errorf("error", er)
		goto new
	}
	if x != 2 || x != 1 {
		fmt.Println("Enter one of them!")
		goto new
	}
	if x == 1 {

	} else {

	}
}

func Sell(name string, quantity uint64) {
	//if data.Sell(name, quantity) {
	//
	//} else {
	//	fmt.Println("We dont have this product!")
	//}
}

func (i Inventory) Delete() {

}
