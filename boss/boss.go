package boss

import "fmt"

type Boss struct {
	BossName     string
	BossPassword string
	Budget       int
	Profit       int
}

func NewBoss(budget int, profit int) Boss {
	return Boss{
		BossName:     "Abdurahmonjon",
		BossPassword: "@bdurahmojon2002",
		Budget:       budget,
		Profit:       profit,
	}
}
func Bossmain() {
	b := NewBoss(100000, 0)
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
		fmt.Println("			1-Viewing products list" +
			"			2-Add or increase products" +
			"			3-Delete expired products" +
			"			4-Exit to main")
		_, err := fmt.Scan(&boss)
		if err != nil {
			fmt.Println("Error while entering choice")
		}
		if boss > 4 && boss < 1 {
			fmt.Println("Invailed choice!")
		} else {
			if boss == 1 {
				
			}
		}
	}
}
