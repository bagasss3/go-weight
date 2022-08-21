package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var (
	LayoutISO              string        = "2006-01-02"
	MaxIdleConnsDefault    int           = 2
	MaxOpenConnsDefault    int           = 20
	ConnMaxLifeTimeDefault time.Duration = 1 * time.Hour
)

// DatabaseDSN :nodoc:
func DatabaseDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SSLMODE"))
}

func PostgreMaxIdleConns() int {
	MaxIdleConns := os.Getenv("MAX_CONN_POOL")
	if MaxIdleConns == "" {
		return MaxIdleConnsDefault
	}
	MaxIdleConnsInt, err := strconv.Atoi(MaxIdleConns)
	if err != nil {
		log.Print("Converting MaxIdleConn failed, using default value instead")
		return MaxIdleConnsDefault
	}
	return MaxIdleConnsInt
}

func PostgreMaxOpenConns() int {
	MaxOpenConns := os.Getenv("MAX_OPEN_CONN")
	if MaxOpenConns == "" {
		return MaxOpenConnsDefault
	}
	MaxOpenConnsInt, err := strconv.Atoi(MaxOpenConns)
	if err != nil {
		log.Print("Converting MaxOpenConns failed, using default value instead")
		return MaxOpenConnsDefault
	}
	return MaxOpenConnsInt
}

func PostgreSetConnMaxLifeTime() time.Duration {
	ConnMaxLifeTime := os.Getenv("MAX_LIFE_TIME")
	if ConnMaxLifeTime == "" {
		return ConnMaxLifeTimeDefault
	}
	ConnMaxLifeTimeDur, err := time.ParseDuration(ConnMaxLifeTime)
	if err != nil {
		return ConnMaxLifeTimeDefault
	}
	return ConnMaxLifeTimeDur
}
