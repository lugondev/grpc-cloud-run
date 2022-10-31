package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	db "waas-service/db/sqlc"
)

type DBConfig struct {
	Host     string
	Port     int
	DBName   string
	Username string
	Password string
	SSLMode  string
	TimeZone string

	ConnectionMaxIdleTime int
	ConnectionMaxLifetime int
	MaxIdleConnections    int
	MaxOpenConnections    int

	Logger struct {
		LogLevel string
	}
}

func NewDB() (*db.SQLStore, error) {
	var config DBConfig
	if err := viper.UnmarshalKey("db.postgresql.primary", &config); err != nil {
		panic(err)
	}

	dsn := getDSN(config)
	openDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db.NewStore(openDB), nil
}

func getDSN(configuration DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s TimeZone=%s",
		configuration.Host,
		configuration.Port,
		configuration.DBName,
		configuration.Username,
		configuration.Password,
		configuration.SSLMode,
		configuration.TimeZone,
	)
}
