package main

// Item contiene información de la factura
type Item struct {
	id uint
	product string
	value float64
}

func Mew(id uint, product string, value float64) Iten {
	return Item{id, product, value}
}