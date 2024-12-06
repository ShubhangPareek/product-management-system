package utils

import "github.com/gin-gonic/gin"

func HandleError(c *gin.Context, statusCode int, errMessage string) {
    c.JSON(statusCode, gin.H{
        "error": errMessage,
    })
}