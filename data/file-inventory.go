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
	js, err := json.MarshalIndent(list, "", "	")
	if err != nil {
		return false, err
	}
	err = os.WriteFile("data/data.json", js, 0600)
	if err != nil {
		return false, err
	}
	return true, nil
}
func Sell() {

}
