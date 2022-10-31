package invoiceitem

// Item contiene información de la factura
type Item struct {
	id uint
	product string
	value float64
}

// New retorna un nuevo item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// getter de itemvalue
func (i Item) Value() float64 {
	return i.value
}