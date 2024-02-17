package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	MySQLHost     string
	MySQLUser     string
	MySQLPassword string
	MySQLDBName   string
	MySQLPort     string
)

func LoadEnvVars() {
	dir, _ := os.Getwd()
	AppPath := dir

	godotenv.Load(filepath.Join(AppPath, "/.env"))

	MySQLHost = os.Getenv("MYSQL_HOST")
	MySQLUser = os.Getenv("MYSQL_USERNAME")
	MySQLPassword = os.Getenv("MYSQL_PASSWORD")
	MySQLDBName = os.Getenv("MYSQL_DBNAME")
	MySQLPort = os.Getenv("MYSQL_PORT")
}
