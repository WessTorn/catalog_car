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

// @Summary      Delete a car
// @Tags         cars
// @Description  Delete a car from the database by its registration number
// @Accept       json
// @Produce      json
// @Param        id path string true "Registration number of the car to delete"
// @Success      200 {object} map[string]string "{"message": "Car deleted successfully"}"
// @Failure      500 {object} map[string]string "{"error": "Failed to delete car"}"
// @Router       /cars/{id} [delete]
func DeleteCar(c *gin.Context, db *pg.DB) {
	logger.Log.Info("DELETE request received")
	logger.Log.Debug("DELETE request received (DeleteCar): " + os.Getenv("HOST_RELATIVE_PATH") + "/:id")

	regNum := c.Param("id")

	logger.Log.Debug("Request parameter: " + fmt.Sprintf("%+v", regNum))

	var car data.Car
	_, err := db.Model(&car).Where("reg_num = ?", regNum).Delete()

	if err != nil {
		logger.Log.Debug("[error] Failed to delete car: " + fmt.Sprintf("%+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
