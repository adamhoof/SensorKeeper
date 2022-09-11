package database

import mqtt_update "SensorKeeper/pkg/MQTTUpdateStructs"

type DatabaseHandler interface {
	Connect(config *string) (err error)
	InsertWaterLevelUpdate(update *mqtt_update.WaterLevel)
}
