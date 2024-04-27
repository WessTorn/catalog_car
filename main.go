package main

import (
	"github.com/WessTorn/catalog_car/config"
	"github.com/WessTorn/catalog_car/data"
	"github.com/WessTorn/catalog_car/logger"
	"github.com/WessTorn/catalog_car/queries"
	"os"
	//_ "igor/docs"
)

//	@title			Car catalog
//	@version		1.0.0
//	@description	This is an example of a car catalog API.

//	@host		localhost:8080
//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	config.InitConfig()
	logger.InitLogger()

	logger.Log.Info("Start")

	logger.Log.Debug("Config:")
	logger.Log.Debug("DB_ADDRESS=" + os.Getenv("DB_ADDRESS"))
	logger.Log.Debug("DB_USER=" + os.Getenv("DB_USER"))
	logger.Log.Debug("DB_PASSWORD=" + os.Getenv("DB_PASSWORD"))
	logger.Log.Debug("DB_NAME=" + os.Getenv("DB_NAME"))
	logger.Log.Debug("HOST_RELATIVE_PATH=" + os.Getenv("HOST_RELATIVE_PATH"))
	logger.Log.Debug("EXTERNAL_API_URL=" + os.Getenv("EXTERNAL_API_URL"))

	db := data.ConnectDB()
	if db == nil {
		logger.Log.Fatalf("Failed to connect to database")
	}
	defer db.Close()

	logger.Log.Info("Database connected")

	err := data.CreateSchema(db)
	if err != nil {
		logger.Log.Fatalf("Failed to create database: %v", err)
	}

	router := queries.InitRouter(db)
	logger.Log.Info("Routers init")

	router.Run(os.Getenv("HOST_URL"))

	logger.Log.Info("Routers run")
	logger.Log.Debug("Routers run: " + os.Getenv("HOST_URL"))
}
