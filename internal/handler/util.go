package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func errResponse(c *gin.Context, code int, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}
