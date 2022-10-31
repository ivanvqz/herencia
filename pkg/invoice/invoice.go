package invoice

import (
	"github.com/ivanvqz/herencia/pkg/customer/customer"
	"github.com/ivanvqz/herencia/pkg/invoiceitem/invoiceitem"

)

type Invoice struct {
	Country string
	City   string
	Amount float64
	client customer.Customer // referencia a estructura cliente
	item []invoiceitem.Item //slice de items
}

// New retorna una nueva factura
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice{
	return Invoice {
		country: Country,
		city: City,
		client: client,
		items: items,
	}
}
//setter de total invoice
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.Amount += item.Value()
	}
}	