package mqtt_handlers

import (
	database "SensorKeeper/pkg/Database"
	mqtt_update "SensorKeeper/pkg/MQTTUpdateStructs"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ReceiveFlowerWaterLevelUpdate(dbHandler database.DatabaseHandler) mqtt.MessageHandler {
	return func(mqttClient mqtt.Client, message mqtt.Message) {
		update := mqtt_update.WaterLevel{}
		err := json.Unmarshal(message.Payload(), &update)
		if err != nil {
			fmt.Println("Failed to unpack error into WaterLevelUpdate struct", err)
		}

		dbHandler.InsertWaterLevelUpdate(&update)

		if !update.IsCriticalValue {

		}
	}
}
