package queries

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"igor/data"
	"igor/logger"
	"net/http"
	"os"
	"strconv"
)

// @Summary      Get a list of cars
// @Tags         cars
// @Description  Retrieve a list of cars with optional filtering and pagination
// @Accept       json
// @Produce      json
// @Param        regNum   query     string  false  "Filter by registration number"
// @Param        mark     query     string  false  "Filter by car mark"
// @Param        model    query     string  false  "Filter by car model"
// @Param        year     query     string  false  "Filter by car year"
// @Param        page     query     int     false  "Page number" default(1)
// @Param        limit    query     int     false  "Number of items per page" (10)
// @Success      200  {array}   data.Car
// @Failure      500  {object}   map[string]string "{"error": "Failed to get cars"}"
// @Router       /cars [get]
func GetCars(c *gin.Context, db *pg.DB) {
	logger.Log.Info("GET request received")
	logger.Log.Debug("GET request received (GetCars): " + os.Getenv("HOST_RELATIVE_PATH"))

	var cars []data.Car

	query := db.Model(&cars).Relation("Owner")

	// Фильтрация и пагинация
	regNum := c.Query("regNum")
	if regNum != "" {
		query = query.Where("reg_num = ?", regNum)
	}

	mark := c.Query("mark")
	if mark != "" {
		query = query.Where("mark = ?", mark)
	}

	model := c.Query("model")
	if model != "" {
		query = query.Where("model = ?", model)
	}

	year := c.Query("year")
	if year != "" {
		query = query.Where("year = ?", year)
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		logger.Log.Debug("[error] Failed to page: " + fmt.Sprintf("%+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to page"})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err != nil {
		logger.Log.Debug("[error] Failed to limit: " + fmt.Sprintf("%+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to limit"})
		return
	}

	err = query.Order("id ASC").Offset((page - 1) * limit).Limit(limit).Select()

	if err != nil {
		logger.Log.Debug("[error] Failed to get cars: " + fmt.Sprintf("%+v", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cars"})
		return
	}

	for i := range cars {
		if cars[i].Owner == nil {
			var owner data.Owner
			err = db.Model(&owner).Where("id = ?", cars[i].OwnerId).First()
			if err == nil {
				cars[i].Owner = &owner
			}
		}
	}

	logger.Log.Debug("Reply to request: " + fmt.Sprintf("%+v", cars))

	c.JSON(http.StatusOK, cars)
}
