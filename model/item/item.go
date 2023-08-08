package item

type Item struct {
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	Value       float64 `json:"Value"`
}

// empty array of Items
func EmptyItems() *[]Item {
	return &[]Item{}
}
