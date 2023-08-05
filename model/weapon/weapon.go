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

func NewWeapon(name string, powerRange, maxDamageTake, minDamageTake float64) Weapon {
	return Weapon{
		Name:          name,
		PowerRange:    powerRange,
		MaxDamageTake: maxDamageTake,
		MinDamageTake: minDamageTake,
	}
}
