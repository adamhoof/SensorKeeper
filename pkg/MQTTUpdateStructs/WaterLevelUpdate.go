package mqtt_update

type WaterLevel struct {
	ClientName           string `json:"ClientName"`
	ClientSubscribeTopic string `json:"ClientSubscribeTopic"`
	SensorValue          string `json:"SensorValue"`
	IsCriticalValue      bool   `json:"IsCriticalValue"`
}
