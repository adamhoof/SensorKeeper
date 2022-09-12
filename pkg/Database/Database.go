package database

import mqtt_update "SensorKeeper/pkg/MQTTUpdateStructs"

type DatabaseHandler interface {
	Connect(config *string) (err error)
	InsertSoilMoistureUpdate(update *mqtt_update.SoilMoisture)
}
