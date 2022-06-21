package basket

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type customer struct {
	Name string
	Quantity uint64
	Price uint64
}

type Trash []customer

func (t *Trash) Append(name string, quantity, price uint64)  {
	*t = append(*t, customer{
		Name: name,
		Quantity: quantity,
		Price: price,
	})
}

func (t Trash) Show()  {
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Your basket")
	fmt.Fprintln(w, "%#\t\t\t\tname\t\t\t\tquantity\t\t\t\t")
	for i, p := range t {
		fmt.Fprintf(w, "%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t", i+1, p.Name, p.Quantity)
	}
	w.Flush()
}