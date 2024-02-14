package config

import "os"

var (
	MySQLHost     = os.Getenv("MYSQL_HOST")
	MySQLUser     = os.Getenv("MYSQL_USERNAME")
	MySQLPassword = os.Getenv("MYSQL_PASSWORD")
	MySQLDBName   = os.Getenv("MYSQL_DBNAME")
	MySQLPort     = os.Getenv("MYSQL_PORT")
)
