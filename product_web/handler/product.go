package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProductListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
