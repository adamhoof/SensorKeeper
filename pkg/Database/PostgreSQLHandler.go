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

func (handler *PostgresDBHandler) InsertSoilMoistureUpdate(update *mqtt_update.SoilMoistureUpdate) {
	_, err := handler.db.Exec("INSERT INTO plant_soil_moisture_values (plant_name, value, is_critical_value, time_stamp) VALUES($1, $2, $3, $4);", update.PlantName, update.Value, update.IsCriticalValue, time.Now())
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
