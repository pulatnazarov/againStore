package data

import (
	"encoding/json"
	"github.com/againStore/product"
	"os"
)

func List() product.List {
	file, err := os.ReadFile("data/data.json")
	if err != nil {
		panic(err)
	}
	var list product.List
	err = json.Unmarshal(file, &list)
	if err != nil {
		panic(err)
	}
	return list
}

func Add(name string, n, pr, op uint64) (bool, error) {
	list := List()
	s := 0
	for i, p := range list {
		if p.Name == name {
			list[i].Quantity += n
			s++
		}
	}
	if s == 0 {
		list = append(list, product.Product{Name: name, Quantity: n, Price: pr, OriginPrice: op})
	}
	write(list)
	return true, nil
}
func write(list product.List) {
	js, err := json.MarshalIndent(list, "", "	")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("data/data.json", js, 0600)
	if err != nil {
		panic(err)
	}
}
func Sell(name string, q uint64) (uint64, bool) {
	list := List()
	for i, p := range list {
		if p.Name == name {
			if p.Quantity-q > 0 {
				list[i].Quantity -= q
				write(list)
				return p.Price, true
			} else if p.Quantity-q < 0 {
				return 0, false
			}
			del(i, &list)
			write(list)
			return p.Price, true
		}
	}
	return 0, false
}
func del(i int, l *product.List) {
	*l = append((*l)[:i], (*l)[i+1:]...)
}
func Dell(name string) {
	list := List()
	for i, p := range list {
		if p.Name == name {
			del(i, &list)
			write(list)
		}
	}
}
