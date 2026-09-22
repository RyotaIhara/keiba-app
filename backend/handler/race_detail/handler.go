// Package racedetail
package racedetail

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	racingModel "keiba-app-backend/model/racing"
	raceService "keiba-app-backend/service/racing"
)

type service interface {
	GetRaceDetails(raceID int64) ([]racingModel.RaceDetail, error)
	GetRaceDetail(raceID, detailID int64) (racingModel.RaceDetail, error)
}

func parsePositiveID(value, name string, c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": name + " must be a positive integer"})
		return 0, false
	}
	return id, true
}

func Index(service *raceService.RaceDetailService) gin.HandlerFunc {
	return func(c *gin.Context) {
		raceID, ok := parsePositiveID(c.Param("race_id"), "race_id", c)
		if !ok {
			return
		}

		response, err := service.GetRaceDetails(raceID)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch race details"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}

func Show(service service) gin.HandlerFunc {
	return func(c *gin.Context) {
		raceID, ok := parsePositiveID(c.Param("race_id"), "race_id", c)
		if !ok {
			return
		}
		detailID, ok := parsePositiveID(c.Param("race_detail_id"), "race_detail_id", c)
		if !ok {
			return
		}

		response, err := service.GetRaceDetail(raceID, detailID)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "race detail not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch race detail"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}
