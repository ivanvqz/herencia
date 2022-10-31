package invoice
import "module github.com/ivanvqz/herencia/pkg/customer"

type Invoice struct {
	Country string
	City   string
	Amount float64
	client customer.Customer // referencia a estructura cliente
	items invoiceitem.
}