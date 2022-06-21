package boss

import (
	"fmt"
	"github.com/againStore/data"
	"github.com/againStore/inventory"
	"os"
	"text/tabwriter"
)

var BassFile Boss

type Boss struct {
	BossName     string
	BossPassword string
	Budget       uint64
	Profit       uint64
}

func NewBoss(name, password string, budget uint64) Boss {
	return Boss{
		BossName:     name,
		BossPassword: password,
		Budget:       budget,
		Profit:       0,
	}
}
func MainBoss(b Boss) {
	var askPassword string
	chance := 4
	for chance > 0 {
		fmt.Print("Enter the password: ")
		_, err := fmt.Scan(&askPassword)
		if err != nil {
			fmt.Print("Error while asking password: ", err)
		}
		if askPassword == b.BossPassword {
			fmt.Print("				Welcome our boss")
			break
		} else {
			chance -= 1
			fmt.Println("Incorrect password!!!")
		}
		if chance == 0 {
			fmt.Println("Your chanses expired <boss> :(   ")
		}
	}
	for {
		var boss int
		fmt.Print(`
		1 - Viewing products list
		2 - Add or increase products
		3 - Delete expired products
		4 - Exit to main
	`)
		_, err := fmt.Scan(&boss)
		if err != nil {
			fmt.Println("Error while entering choice")
		}
		if boss > 4 && boss < 1 {
			fmt.Println("Invailed choice!")
		} else {
			i := inventory.Inventory{}
			if boss == 1 {
				w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', tabwriter.TabIndent)
				fmt.Fprintln(w, "#\t\t\t\tname\t\t\t\tquantity\t\t\t\t")
				for i, p := range data.List() {
					fmt.Fprintf(w, "%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t\n", i+1, p.Name, p.Quantity)
				}
				w.Flush()
			} else if boss == 2 {
				i.AddOrIncrease()
			} else if boss == 3 {
				i.Delete()
			} else if boss == 4 {
				return
			}
		}
	}
}
