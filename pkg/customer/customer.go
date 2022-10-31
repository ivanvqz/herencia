package customer

// customer is a struct that contains information about a customer.
type Customer struct {
	Name    string
	addres   string
	phone  string
}

func new(name, addres, phone string) Customer  {
	return Customer{
		Name:    name,
		addres:   addres,
		phone:  phone,
	}
}