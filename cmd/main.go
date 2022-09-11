package main

import (
	database "SensorKeeper/pkg/Database"
	essential "SensorKeeper/pkg/EssentialConfig"
	mqtt_handlers "SensorKeeper/pkg/MQTTMessageHandlers"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

func main() {
	var conf essential.Config
	essential.LoadConfigFromJsonFile(os.Args[1], &conf)

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		conf.DbHost,
		conf.DbPort,
		conf.DbUser,
		conf.DbUserPassword,
		conf.DbName)

	var postgresDBHandler = database.PostgresDBHandler{}
	err := postgresDBHandler.Connect(&dbConnectionString)
	if err != nil {
		panic("could not connect to database: " + err.Error())
	}

	options := mqtt.ClientOptions{}
	options.AddBroker(conf.MQTTServer)
	options.SetClientID(conf.MQTTClientName)
	options.SetAutoReconnect(true)
	options.SetConnectRetry(true)
	options.SetCleanSession(false)
	options.SetOrderMatters(false)

	mqttClient := mqtt.NewClient(&options)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("MQTT client connection established?: ", mqttClient.IsConnected())

	mqttClient.Subscribe(conf.MQTTSensorPublishTopic, 0, mqtt_handlers.ReceiveFlowerWaterLevelUpdate(&postgresDBHandler))

}