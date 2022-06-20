package main

import (
	"bufio"
	"fmt"
	"github.com/againStore/product"
	"os"
	"strconv"
	"strings"
)

var List = product.List{}

func main() {
	start()
}

func RightInput() (string, error) {
	sentence, er := bufio.NewReader(os.Stdin).ReadString('\n')
	return sentence, er
}

func ReadFromFile() product.List {
	p := product.Product{}
	f, er := os.OpenFile("data/data.txt", os.O_RDWR, 0600)
	if er != nil {
		fmt.Println("Error while opening file", er)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		list := strings.Fields(scanner.Text())
		if len(list) < 4 {
			fmt.Println("Error while reading line:", er)
			return nil
		}
		p.Name = list[0]
		p.Quantity, p.Price, p.OriginPrice = converting(list)
		List = append(List, p)
	}
	return List
}

func converting(list []string) (uint64, uint64, uint64) {
	a, er := strconv.ParseUint(list[1], 10, 0)
	if er != nil {
		fmt.Println("Error while converting to uint", er)
	}
	b, er := strconv.ParseUint(list[2], 10, 0)
	if er != nil {
		fmt.Println("Error while converting to uint", er)
	}
	c, er := strconv.ParseUint(list[3], 10, 0)
	if er != nil {
		fmt.Println("Error while converting to uint", er)
	}
	return a, b, c
}

func start() {
	//List := ReadFromFile()
	for {
	end:
		fmt.Print(`
		1 - Start
		2 - Exit
	`)
		sentence, er := RightInput()
		if er != nil {
			fmt.Errorf("error: %d", er)
			goto end
		}
		x, er := strconv.Atoi(sentence)
		if er != nil {
			fmt.Printf("error: %d", er)
			goto end
		}
		if x > 2 || x < 1 {
			fmt.Println("Enter one of them")
			goto end
		}
		if x == 1 {
			//Sell(List)
		} else if x == 2 {
			break
		}
	}
}
