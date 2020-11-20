package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": "invalid request body",
	})
}

func NotFound(c *gin.Context, id int) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":   true,
		"message": "the id=" + strconv.Itoa(id) + " does not exist",
	})
}
