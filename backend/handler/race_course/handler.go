// Package racecourse
package racecourse

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	handlerHelper "keiba-app-backend/helper"
	raceCourseModel "keiba-app-backend/model/race_course"
)

type raceCourseService interface {
	GetRaceCourses() ([]raceCourseModel.Racecourse, error)
	GetRaceCourse(id int64) (raceCourseModel.Racecourse, error)
	CreateRaceCourse(code, name string) (raceCourseModel.Racecourse, error)
	UpdateRaceCourse(id int64, code, name string) (raceCourseModel.Racecourse, error)
	DeleteRaceCourse(id int64) error
}

type request struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// Index 競馬場一覧を取得するハンドラ
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

// Show 指定されたIDの競馬場を取得するハンドラ
func Show(service raceCourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := handlerHelper.ParseID("", c)
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

// Create 競馬場を作成するハンドラ
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

// Update 指定されたIDの競馬場を更新するハンドラ
func Update(service raceCourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := handlerHelper.ParseID("", c)
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

// Delete 指定されたIDの競馬場を削除するハンドラ
func Delete(service raceCourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := handlerHelper.ParseID("", c)
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
