// Package racecourse
package racecourse

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	raceCourseModel "keiba-app-backend/model/race_course"
)

type raceCourseService interface {
	GetRaceCourses() ([]raceCourseModel.Racecourse, error)
	GetRaceCourse(id int64) (raceCourseModel.Racecourse, error)
	CreateRaceCourse(code, name string) (raceCourseModel.Racecourse, error)
	UpdateRaceCourse(id int64, code, name string) (raceCourseModel.Racecourse, error)
	DeleteRaceCourse(id int64) error
}

func Index(service raceCourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := service.GetRaceCourses()
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch race courses"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}

type request struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}
	return id, true
}

func Show(service raceCourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseID(c)
		if !ok {
			return
		}
		response, err := service.GetRaceCourse(id)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "race course not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch race course"})
			return
		}
		c.IndentedJSON(http.StatusOK, response)
	}
}

func Create(service raceCourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := service.CreateRaceCourse(req.Code, req.Name)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create race course"})
			return
		}
		c.IndentedJSON(http.StatusCreated, response)
	}
}

func Update(service raceCourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseID(c)
		if !ok {
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := service.UpdateRaceCourse(id, req.Code, req.Name)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "race course not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update race course"})
			return
		}
		c.IndentedJSON(http.StatusOK, response)
	}
}

func Delete(service raceCourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseID(c)
		if !ok {
			return
		}
		err := service.DeleteRaceCourse(id)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "race course not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete race course"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
