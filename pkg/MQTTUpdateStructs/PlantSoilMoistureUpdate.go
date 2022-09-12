package mqtt_update

type SoilMoistureUpdate struct {
	PlantName       string `json:"PlantName"`
	Value           uint16 `json:"Value"`
	IsCriticalValue bool   `json:"IsCriticalValue"`
}
