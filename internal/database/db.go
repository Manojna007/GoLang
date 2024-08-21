package database

import (
	"fmt"
	"time"

	"GoAssignment/internal/config"
	"GoAssignment/internal/logger"

	_ "github.com/denisenkom/go-mssqldb" // MS SQL Server driver
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() error {
	// Connection string for SQL Server with SQL Server Authentication
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		config.AppConfig.DB.User,
		config.AppConfig.DB.Password,
		config.AppConfig.DB.Host,
		config.AppConfig.DB.Port,
		config.AppConfig.DB.Name,
	)

	var err error
	DB, err = sqlx.Open("sqlserver", connStr)
	if err != nil {
		logger.Error("Failed to open database connection:", err)
		return err
	}

	// Ensure the database is reachable
	if err := DB.Ping(); err != nil {
		logger.Error("Database ping failed:", err)
		return err
	}

	logger.Info("Database connection established successfully")
	return nil
}

// CheckDB continuously pings the database to check its availability
func CheckDB(healthChan chan<- bool) {
	defer close(healthChan)

	for {
		select {
		case <-time.After(10 * time.Second):
			err := DB.Ping()
			if err != nil {
				healthChan <- false
				logger.Error("Database connection error:", err)
			} else {
				healthChan <- true
			}
		}
	}
}
