// Package user
package user

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	handlerHelper "keiba-app-backend/helper"
	userModel "keiba-app-backend/model/user"
)

type userService interface {
	GetUsers() ([]userModel.User, error)
	GetUser(id int64) (userModel.User, error)
	CreateUser(code, name, password string) (userModel.User, error)
	UpdateUser(id int64, code, name string) (userModel.User, error)
	DeleteUser(id int64) error
}

type createRequest struct {
	Code     string `json:"code" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type updateRequest struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// Index ユーザー一覧を取得するハンドラ
func Index(service userService) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := service.GetUsers()
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}

// Show 指定されたIDのユーザーを取得するハンドラ
func Show(service userService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := handlerHelper.ParseID("", c)
		if !ok {
			return
		}

		response, err := service.GetUser(id)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}

// Create ユーザーを作成するハンドラ
func Create(service userService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request createRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := service.CreateUser(request.Code, request.Name, request.Password)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

		c.IndentedJSON(http.StatusCreated, response)
	}
}

// Update 指定されたIDのユーザーを更新するハンドラ
func Update(service userService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := handlerHelper.ParseID("", c)
		if !ok {
			return
		}

		var request updateRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := service.UpdateUser(id, request.Code, request.Name)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}

// Delete 指定されたIDのユーザーを削除するハンドラ
func Delete(service userService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := handlerHelper.ParseID("", c)
		if !ok {
			return
		}

		err := service.DeleteUser(id)
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
