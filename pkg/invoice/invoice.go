package invoice

type Invoice struct {
	Country string
	City   string
	Amount float64
	client customer.Customer // referencia a estructura cliente
}