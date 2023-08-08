package item

type Item struct {
	Key         string  `json:"Key"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	Value       float64 `json:"Value"`
}

// empty array of Items
func EmptyItems() *[]Item {
	return &[]Item{}
}
