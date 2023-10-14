package util

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pingcap/log"
)

var errMissingEnv = errors.New("failed to load env vars")

// DBConfig holds the database config needed to connect
type DBConfig struct {
	DBDriver string
	DBSource string
}

const projectDirName = "simplebank"

// LoadEnv loads the environment variables set on the .env file
func loadEnv() error {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return godotenv.Load(string(rootPath) + `/.env`)
}

// GetDBConfig returns the database config data
func GetDBConfig() (*DBConfig, error) {
	err := loadEnv()
	if err != nil {
		log.S().Warnf("failed to load envfile: %s", err.Error())
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbDriver = strings.TrimSpace(dbDriver)
	if len(dbDriver) == 0 {
		return nil, errMissingEnv
	}

	dbUser := os.Getenv("DB_USER")
	dbUser = strings.TrimSpace(dbUser)
	if len(dbUser) == 0 {
		return nil, errMissingEnv
	}

	dbPass := os.Getenv("DB_PASS")
	dbPass = strings.TrimSpace(dbPass)
	if len(dbPass) == 0 {
		return nil, errMissingEnv
	}

	dbName := os.Getenv("DB_NAME")
	dbName = strings.TrimSpace(dbName)
	if len(dbName) == 0 {
		return nil, errMissingEnv
	}

	dbHost := os.Getenv("DB_HOST")
	dbHost = strings.TrimSpace(dbHost)
	if len(dbHost) == 0 {
		return nil, errMissingEnv
	}

	dbPort := os.Getenv("DB_PORT")
	dbPort = strings.TrimSpace(dbPort)
	if len(dbPort) == 0 {
		return nil, errMissingEnv
	}

	dbSource := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbDriver, dbUser, dbPass, dbHost, dbPort, dbName)
	return &DBConfig{
		DBDriver: dbDriver,
		DBSource: dbSource,
	}, nil
}
