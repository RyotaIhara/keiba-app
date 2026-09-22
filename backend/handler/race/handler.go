// Package race
package race

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	racingSupport "keiba-app-backend/model/racing/support"
	raceTypes "keiba-app-backend/model/racing/types/race"
	raceService "keiba-app-backend/service/racing"
)

func Index(service *raceService.RaceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := service.GetRaces()
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch races"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}

type raceRequest struct {
	RaceDate       string                   `json:"race_date" binding:"required"`
	RaceCourseID   int64                    `json:"race_course_id" binding:"required"`
	RaceNumber     int                      `json:"race_number" binding:"required"`
	RaceName       string                   `json:"race_name" binding:"required"`
	StartTime      string                   `json:"start_time" binding:"required"`
	Surface        raceTypes.Surface        `json:"surface" binding:"required"`
	Distance       int                      `json:"distance" binding:"required"`
	Direction      raceTypes.Direction      `json:"direction" binding:"required"`
	Weather        raceTypes.Weather        `json:"weather" binding:"required"`
	TrackCondition raceTypes.TrackCondition `json:"track_condition" binding:"required"`
	RaceConditions string                   `json:"race_conditions" binding:"required"`
}

func parseRaceRequest(request raceRequest) (racingSupport.RaceInput, error) {
	raceDate, err := time.Parse("2006-01-02", request.RaceDate)
	if err != nil {
		return racingSupport.RaceInput{}, errors.New("race_date must use YYYY-MM-DD")
	}

	startTime, err := time.Parse("15:04:05", request.StartTime)
	if err != nil {
		return racingSupport.RaceInput{}, errors.New("start_time must use HH:MM:SS")
	}

	if request.RaceCourseID <= 0 || request.RaceNumber <= 0 || request.Distance <= 0 {
		return racingSupport.RaceInput{}, errors.New("race_course_id, race_number, and distance must be positive")
	}
	if request.Surface != raceTypes.SurfaceTurf && request.Surface != raceTypes.SurfaceDirt {
		return racingSupport.RaceInput{}, errors.New("invalid surface")
	}
	if request.Direction != raceTypes.DirectionRight && request.Direction != raceTypes.DirectionLeft {
		return racingSupport.RaceInput{}, errors.New("invalid direction")
	}
	if request.Weather != raceTypes.WeatherSunny &&
		request.Weather != raceTypes.WeatherCloudy &&
		request.Weather != raceTypes.WeatherRainy &&
		request.Weather != raceTypes.WeatherSnowy {
		return racingSupport.RaceInput{}, errors.New("invalid weather")
	}
	if request.TrackCondition != raceTypes.TrackConditionFirm &&
		request.TrackCondition != raceTypes.TrackConditionGood &&
		request.TrackCondition != raceTypes.TrackConditionYield &&
		request.TrackCondition != raceTypes.TrackConditionSoft {
		return racingSupport.RaceInput{}, errors.New("invalid track_condition")
	}

	return racingSupport.RaceInput{
		RaceDate:       raceDate,
		RaceCourseID:   request.RaceCourseID,
		RaceNumber:     request.RaceNumber,
		RaceName:       request.RaceName,
		StartTime:      startTime,
		Surface:        request.Surface,
		Distance:       request.Distance,
		Direction:      request.Direction,
		Weather:        request.Weather,
		TrackCondition: request.TrackCondition,
		RaceConditions: request.RaceConditions,
	}, nil
}

func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}

func bindRaceRequest(c *gin.Context) (racingSupport.RaceInput, bool) {
	var request raceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return racingSupport.RaceInput{}, false
	}

	input, err := parseRaceRequest(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return racingSupport.RaceInput{}, false
	}
	return input, true
}

func Show(service *raceService.RaceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseID(c)
		if !ok {
			return
		}

		response, err := service.GetRace(id)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "race not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch race"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}

func Create(service *raceService.RaceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		input, ok := bindRaceRequest(c)
		if !ok {
			return
		}

		response, err := service.CreateRace(input)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create race"})
			return
		}

		c.IndentedJSON(http.StatusCreated, response)
	}
}

func Update(service *raceService.RaceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseID(c)
		if !ok {
			return
		}
		input, ok := bindRaceRequest(c)
		if !ok {
			return
		}

		response, err := service.UpdateRace(id, input)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "race not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update race"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}

func Delete(service *raceService.RaceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseID(c)
		if !ok {
			return
		}

		err := service.DeleteRace(id)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "race not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete race"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
