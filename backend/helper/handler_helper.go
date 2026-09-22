package helper

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ParseID 指定された値、またはパスパラメータからIDを解析する
func ParseID(value string, c *gin.Context) (int64, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		value = c.Param("id")
	}

	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}

	return id, true
}
