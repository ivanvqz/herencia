package invoice

import (
	"github.com/ivanvqz/herencia/pkg/customer"
	"github.com/ivanvqz/herencia/pkg/invoiceitem"

)

type Invoice struct {
	country string
	city   string
	total float64
	client customer.Customer // referencia a estructura cliente
	items []invoiceitem.Item //slice de items
}

// New retorna una nueva factura
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice{
	return Invoice {
		country: country,
		city: city,
		client: client,
		items: items,
	}
}
//setter de total invoice
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}	