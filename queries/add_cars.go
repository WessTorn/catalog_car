package queries

import (
	"fmt"
	"github.com/WessTorn/catalog_car/data"
	"github.com/WessTorn/catalog_car/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"net/http"
	"os"
)

// @Summary Add a new car
// @Tags cars
// @Description Add a new car to the database using registration numbers
// @Accept json
// @Produce json
// @Param regNums body []string true "Array of registration numbers"
// @Success 200 {object} data.Car "Car(s) added successfully"
// @Failure 400 {object} map[string]string "{"error": "Invalid request payload"}"
// @Failure 409 {object} map[string]string "{"error": "Car with this regNum already exists"}"
// @Failure 500 {object} map[string]string "{"error": "Failed to add car or check existing owner"}"
// @Router /cars [post]
func AddCar(c *gin.Context, db *pg.DB) {
	logger.Log.Info("POST request received")
	logger.Log.Debug("POST request received (AddCar): " + os.Getenv("HOST_RELATIVE_PATH"))

	var request struct {
		RegNums []string `json:"regNums"`
	}

	err := c.BindJSON(&request)
	if err != nil {
		logger.Log.Debug("[error] Invalid request payload: " + fmt.Sprintf("%+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	logger.Log.Debug("Request payload: " + fmt.Sprintf("%+v", request))

	for _, regNum := range request.RegNums {
		// Берем данные из внешнего API
		var carData *data.Car
		carData, err = GetCarDataFromExternalAPI(regNum)

		if err != nil {
			logger.Log.Debug("[error] Failed to get car data from external API")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get car data from external API"})
			continue
		}

		// Проверяем если ли уже машина с таким же номером в нашей бд
		err = db.Model(&data.Car{}).Where("reg_num = ?", regNum).First()

		if err == nil {
			logger.Log.Debug("[error] Car with this regNum already exists" + fmt.Sprintf("%+v", carData))
			c.JSON(http.StatusConflict, gin.H{"error": "Car with this regNum already exists"})
			continue
		}

		// Проверяем есть ли владелец у нас в бд.
		var existingOwner data.Owner
		err = db.Model(&existingOwner).Where("name = ? AND surname = ? AND patronymic = ?", carData.Owner.Name, carData.Owner.Surname, carData.Owner.Patronymic).First()

		if err != nil && err != pg.ErrNoRows {
			logger.Log.Debug("[error] Failed to check existing owner: " + fmt.Sprintf("%+v", err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing owner"})
			continue
		}

		if existingOwner.ID != 0 {
			// Если владелец уже существует, берем его
			carData.Owner = &existingOwner

			logger.Log.Debug("Take existing owner: " + fmt.Sprintf("%+v", existingOwner))
		} else {
			// Если владельца нет в базе данных, добавляем его и получаем его идентификатор
			owner := data.Owner{
				Name:       carData.Owner.Name,
				Surname:    carData.Owner.Surname,
				Patronymic: carData.Owner.Patronymic,
			}
			_, err = db.Model(&owner).Insert()

			if err != nil {
				logger.Log.Debug("[error] Failed to add owner: " + fmt.Sprintf("%+v", err))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add owner"})
				continue
			}
			carData.Owner = &owner

			logger.Log.Debug("Create new owner: " + fmt.Sprintf("%+v", owner))
		}

		// Добавляем данные в новую машину
		newCar := data.Car{
			RegNum:  regNum,
			Mark:    carData.Mark,
			Model:   carData.Model,
			Year:    carData.Year,
			Owner:   carData.Owner,
			OwnerId: carData.Owner.ID,
		}
		logger.Log.Debug("Create new car: " + fmt.Sprintf("%+v", newCar))

		_, err = db.Model(&newCar).Insert()

		if err != nil {
			logger.Log.Debug("[error] Failed to add car: " + fmt.Sprintf("%+v", err))

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add car"})
			continue
		}

		c.IndentedJSON(http.StatusCreated, newCar)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car(s) added successfully"})
}
