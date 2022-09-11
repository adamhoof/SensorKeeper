package database

import (
	essential "SensorKeeper/pkg/EssentialConfig"
	"fmt"
	"testing"
	"time"
)

func TestInsertWaterLevelUpdate(t *testing.T) {
	var conf essential.Config
	essential.LoadConfigFromJsonFile("/home/adamhoof/Projects/SensorKeeper/ConfigTest2.json", &conf)

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		conf.DbHost,
		conf.DbPort,
		conf.DbUser,
		conf.DbUserPassword,
		conf.DbName)

	var postgresDBHandler = PostgresDBHandler{}
	err := postgresDBHandler.Connect(&dbConnectionString)
	if err != nil {
		t.Errorf("could not connect to database: %s", err)
	}

	_, err = postgresDBHandler.db.Exec("INSERT INTO water_level_sensor_values (sensor_name, value, is_critical, time_stamp) VALUES($1, $2, $3, $4);", "Name", "Value", true, time.Now())
	if err != nil {
		t.Errorf("Failed to insert water level update into db: %s", err)
		err = postgresDBHandler.db.Close()
		return
	}

	err = postgresDBHandler.db.Close()
	if err != nil {
		t.Errorf("Failed to close connection to db: %s", err)
		return
	}

	t.Logf("PASSED")
}
