package database

import (
	mqtt_update "SensorKeeper/pkg/MQTTUpdateStructs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type PostgresDBHandler struct {
	db *sql.DB
}

func (handler *PostgresDBHandler) Connect(config *string) (err error) {
	handler.db, err = sql.Open("postgres", *config)
	if err != nil {
		return fmt.Errorf("could not open connection %s", err)
	}
	return handler.db.Ping()
}

func (handler *PostgresDBHandler) InsertWaterLevelUpdate(update *mqtt_update.WaterLevel) {
	_, err := handler.db.Exec("INSERT INTO water_level_sensor_values (sensor_name, value, is_critical, time_stamp) VALUES($1, $2, $3, $4);", update.ClientName, update.SensorValue, update.IsCriticalValue, time.Now())
	if err != nil {
		fmt.Println("Failed to insert water level update into db", err)
	}
}

func (handler *PostgresDBHandler) ExecuteStatement(statement string) (err error) {
	_, err = handler.db.Exec(statement)
	if err != nil {
		return fmt.Errorf("failed to execute db statement %s", err)
	}
	return err
}
