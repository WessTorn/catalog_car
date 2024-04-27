package queries

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"igor/data"
	"igor/logger"
	"net/http"
	"os"
)

// @Summary		Update car
// @Tags			cars
// @Description	Update a car by registration number
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Registration number of the car"
// @Param			car	body		data.Car	true	"Car data"
// @Success		200	{object} map[string]string "{"message": "Car updated successfully"}"
// @Failure		400	{object} map[string]string "{"error": "Invalid request payload"}"
// @Failure		500	{object} map[string]string "{"error": "Failed to update car"}"
// @Router			/cars/{id} [put]
func UpdateCar(c *gin.Context, db *pg.DB) {
	logger.Log.Info("PUT request received")
	logger.Log.Debug("PUT request received (UpdateCar): " + os.Getenv("HOST_RELATIVE_PATH") + "/:id")

	regNum := c.Param("id")

	logger.Log.Debug("Request parameter: " + fmt.Sprintf("%+v", regNum))

	var updatedCar data.Car
	err := c.BindJSON(&updatedCar)

	if err != nil {
		logger.Log.Debug("[error] Invalid request payload: " + fmt.Sprintf("%+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	logger.Log.Debug("Request payload: " + fmt.Sprintf("%+v", updatedCar))

	updatedCar.RegNum = regNum

	_, err = db.Model(&updatedCar).Where("reg_num = ?", regNum).UpdateNotZero()

	if err != nil {
		logger.Log.Debug("[error] Failed to update car: " + fmt.Sprintf("%+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car updated successfully"})
}
