package essential

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbTableName            string `json:"DbTableName"`
	DbHost                 string `json:"DbHost"`
	DbName                 string `json:"DbName"`
	DbPort                 string `json:"DbPort"`
	DbUser                 string `json:"DbUser"`
	DbUserPassword         string `json:"DbUserPassword"`
	MQTTClientName         string `json:"MQTTClientName"`
	MQTTServer             string `json:"MQTTServer"`
	MQTTServerUsername     string `json:"MQTTServerUsername"`
	MQTTServerPassword     string `json:"MQTTServerPassword"`
	MQTTSensorPublishTopic string `json:"MQTTSensorPublishTopic"`
	TelegramBotToken       string `json:"TelegramBotToken"`
	TelegramBotOwner       string `json:"TelegramBotOwner"`
}

func LoadConfigFromJsonFile(pathToJsonFile string, unpackInto *Config) {
	file, err := os.ReadFile(pathToJsonFile)
	if err != nil {
		panic("Failed to read .json file, make sure you have this file, check spelling, path to file: " + err.Error())
	}

	err = json.Unmarshal(file, &unpackInto)
	if err != nil {
		panic("Failed to extract Config.json into struct: " + err.Error())
	}
}
