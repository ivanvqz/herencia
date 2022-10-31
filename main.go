package main

import (
	"fmt"

	"github.com/ivanvqz/herencia/pkg/invoiceitem"
	"github.com/ivanvqz/herencia/pkg/invoice"
	"github.com/ivanvqz/herencia/pkg/customer"

)

func main() {
	i := invoice.New(
		"México",
		"Puebla",
		customer.New("Ivan", "calle i", "12"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "producto 1", 10),
			invoiceitem.New(2, "producto 2", 20),
			invoiceitem.New(3, "producto 3", 30),
			invoiceitem.New(4, "producto 4", 30),
		},
	)
	i.SetTotal()
	fmt.Printf("Total: %+v", i)
}