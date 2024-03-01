package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormLogger "gorm.io/gorm/logger"
)

func ConnectDatabase() (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(getLoggerLevel(os.Getenv("GORM_LOGGER"))),
	})
	if err != nil {
		return nil, err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = pgsqlDB.Ping()
	if err != nil {
		return nil, err
	}

	maxPoolSize, _ := strconv.Atoi(os.Getenv("DATABASE_POOL_SIZE"))
	maxIdleCons, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNS"))
	connMaxLifeTime, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_CONN_LIFETIME"))
	pgsqlDB.SetMaxOpenConns(maxPoolSize)
	pgsqlDB.SetConnMaxLifetime(time.Duration(connMaxLifeTime * int(time.Second)))
	pgsqlDB.SetMaxIdleConns(maxIdleCons)

	return db, nil
}

// getLoggerLevel return gorm log level setup based from env/config of the app
func getLoggerLevel(v string) gormLogger.LogLevel {
	switch v {
	case "error":
		return gormLogger.Error
	case "warn":
		return gormLogger.Warn
	case "info":
		return gormLogger.Info
	default:
		return gormLogger.Silent
	}
}
