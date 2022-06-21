package store

import "fmt"

type Store struct {
}

func (s Store) Sell() {
	fmt.Printf(`
	1-Shop
	2-List
	3-Exit
	`)

}

// sentence, er := RightInput()
// if er != nil {
// 	fmt.Println(" ", er)
// 	goto end
// }

// sentence = sentence[:len(sentence-1)]
// x, er := strconv.Atoi(sentence)
// if er != nil {
// 	fmt.Println(er)
// 	goto end
// }
// if x>2 || x<1 {
// 	fmt.Println(er)
// }else if x == 1{
// 	return
// }
