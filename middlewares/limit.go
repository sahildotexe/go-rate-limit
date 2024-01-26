package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	clientController "github.com/sahildotexe/go-rate-limit/controllers"
)

func Limit(c *gin.Context) {
	key := c.Request.Header.Get("X-Client-Key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized request please provide X-Client-Key in header"})
		c.Abort()
		return
	}
	client, error := clientController.GetBucket(key)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unauthorized request please provide valid X-Client-Key in header"})
		c.Abort()
		return
	}
	if !client.IsRequestAllowed(1) {
		c.JSON(http.StatusTooManyRequests, gin.H{"message": "Too many requests"})
		c.Abort()
		return
	}
	c.Next()
}
