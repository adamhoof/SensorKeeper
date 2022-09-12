package mqtt_update

type SoilMoisture struct {
	ClientName           string `json:"ClientName"`
	ClientSubscribeTopic string `json:"ClientSubscribeTopic"`
	SensorValue          string `json:"SensorValue"`
	IsCriticalValue      bool   `json:"IsCriticalValue"`
}
