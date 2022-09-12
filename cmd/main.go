package main

import (
	database "SensorKeeper/pkg/Database"
	essential "SensorKeeper/pkg/EssentialConfig"
	mqtt_handlers "SensorKeeper/pkg/MQTTMessageHandlers"
	telegram "SensorKeeper/pkg/Telegram"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"sync"
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
	fmt.Println("Connected to database...")

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
	fmt.Println("Connected to mqtt broker...")

	boi := telegram.User{Id: conf.TelegramBotOwner}
	botHandler := telegram.BotHandler{Owner: boi}
	botHandler.SetToken(conf.TelegramBotToken)

	mqttClient.Subscribe(conf.MQTTPlantSoilMoistureTopic, 0, mqtt_handlers.ReceivePlantSoilMoistureUpdate(&postgresDBHandler, &botHandler))

	go botHandler.StartBot()
	fmt.Println("Telegram bot running...")
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
