package basket

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type customer struct {
	Name     string
	Quantity uint64
	Price    uint64
}

type Trash []customer

func (t *Trash) Append(name string, quantity, price uint64) {
	*t = append(*t, customer{
		Name:     name,
		Quantity: quantity,
		Price:    price,
	})
}

func (t *Trash) Show() {
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Your basket")
	fmt.Fprintln(w, "#\t\t\t\tname\t\t\t\tquantity\t\t\t\t")
	for i, p := range *t {
		fmt.Fprintf(w, "%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t", i+1, p.Name, p.Quantity)
	}
	w.Flush()
}

func (t *Trash) Check() {
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', tabwriter.TabIndent)
	sum := uint64(0)
	fmt.Fprintln(w, "Your check")
	fmt.Fprintln(w, "%#\t\t\t\tname\t\t\t\tquantity\t\t\t\tprice")
	for i, p := range *t {
		fmt.Fprintf(w, "%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t%d", i+1, p.Name, p.Quantity, p.Price)
		sum += p.Price * p.Quantity
	}
	w.Flush()
	fmt.Println("total sum:", sum)
}
