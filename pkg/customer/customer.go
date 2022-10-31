package customer

// customer is a struct that contains information about a customer.
type Customer struct {
	Name string
	addres string
	phone string
}

// return regresa un nuevo cliente
func New(name, addres, phone string) Customer  {
	return Customer{name, addres, phone}	
}