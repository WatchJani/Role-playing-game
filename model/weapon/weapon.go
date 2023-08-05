package weapon

type Weapon struct {
	Name          string  `json:"Name"`
	PowerRange    float64 `json:"PowerRange"`
	MaxDamageTake float64 `json:"MaxDamageTake"`
	MinDamageTake float64 `json:"MinDamageTake"`
}

func EmptyWeapon() *Weapon {
	return &Weapon{}
}
