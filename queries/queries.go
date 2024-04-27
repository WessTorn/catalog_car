package queries

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	//"github.com/swaggo/files"
	//"github.com/swaggo/gin-swagger"
	"github.com/WessTorn/catalog_car/data"
	"github.com/WessTorn/catalog_car/logger"
	"net/http"
	"os"
)

func InitRouter(db *pg.DB) *gin.Engine {
	router := gin.Default()

	router.GET(os.Getenv("HOST_RELATIVE_PATH"), func(c *gin.Context) {
		GetCars(c, db)
	})

	logger.Log.Debug("Init GET router: " + os.Getenv("HOST_RELATIVE_PATH"))

	router.DELETE(os.Getenv("HOST_RELATIVE_PATH")+"/:id", func(c *gin.Context) {
		DeleteCar(c, db)
	})

	logger.Log.Debug("Init DELETE router: " + os.Getenv("HOST_RELATIVE_PATH") + "/:id")

	router.PUT(os.Getenv("HOST_RELATIVE_PATH")+"/:id", func(c *gin.Context) {
		UpdateCar(c, db)
	})

	logger.Log.Debug("Init DELETE router: " + os.Getenv("HOST_RELATIVE_PATH") + "/:id")

	router.POST(os.Getenv("HOST_RELATIVE_PATH"), func(c *gin.Context) {
		AddCar(c, db)
	})

	logger.Log.Debug("Init DELETE router: " + os.Getenv("HOST_RELATIVE_PATH"))

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func GetCarDataFromExternalAPI(regNum string) (*data.Car, error) {
	logger.Log.Info("Getting car data from an external API")
	logger.Log.Debug("Getting car data from an external API: " + os.Getenv("EXTERNAL_API_URL") + "?regNum=" + fmt.Sprintf("%+v", regNum))

	url := os.Getenv("EXTERNAL_API_URL") + "?regNum=" + regNum
	resp, err := http.Get(url)

	if err != nil {
		logger.Log.Debug("Failed to make GET request: " + fmt.Sprintf("%+v", err))
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Log.Debug("Unexpected status code: " + fmt.Sprintf("%d", resp.StatusCode))
		return nil, err
	}

	var carData data.Car
	err = json.NewDecoder(resp.Body).Decode(&carData)

	if err != nil {
		logger.Log.Debug("Failed to decode response:" + fmt.Sprintf("%+v", err))
		return nil, err
	}

	logger.Log.Debug("Getting response from an external API: " + fmt.Sprintf("%+v", carData))

	return &carData, nil
}
