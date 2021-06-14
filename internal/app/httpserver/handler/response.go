package handler

import "github.com/gin-gonic/gin"

func respond(c *gin.Context, statusCode int, obj interface{}) {
	if err, ok := obj.(error); ok {
		c.AbortWithStatusJSON(statusCode, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(statusCode, obj)
}

func respondError(c *gin.Context, statusCode int, err error) {
	respond(c, statusCode, err)
}
