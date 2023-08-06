package item

type Item struct {
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	Value       float64 `json:"Value"`
}

func EmptyItems() *[]Item {
	return &[]Item{}
}

func New(name, description string, value float64) *Item {
	return &Item{
		Name:        name,
		Description: description,
		Value:       value,
	}
}
