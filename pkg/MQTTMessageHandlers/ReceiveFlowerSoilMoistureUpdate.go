package mqtt_handlers

import (
	database "SensorKeeper/pkg/Database"
	mqtt_update "SensorKeeper/pkg/MQTTUpdateStructs"
	telegram "SensorKeeper/pkg/Telegram"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ReceivePlantSoilMoistureUpdate(dbHandler database.DatabaseHandler, botHandler *telegram.BotHandler) mqtt.MessageHandler {
	return func(mqttClient mqtt.Client, message mqtt.Message) {
		update := mqtt_update.SoilMoistureUpdate{}
		err := json.Unmarshal(message.Payload(), &update)
		if err != nil {
			fmt.Println("Failed to unpack error into WaterLevelUpdate struct: ", err)
		}

		dbHandler.InsertSoilMoistureUpdate(&update)

		if !update.IsCriticalValue {
			return
		}
		botHandler.SendText(fmt.Sprintf("Soil moisture critical for plant:\n%s: %d", update.PlantName, update.Value))
	}
}
