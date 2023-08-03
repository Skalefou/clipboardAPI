/*
* name : middlewareClipboard.go
* package : security
* Middleware, a set of processes that is executed before the request is executed to record logs, verify
* if the request is coherent, anti DDOS work...
 */

package security

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddlewareClipboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		var typeRequest int
		if c.Request.Method == "POST" && c.Request.RequestURI == "/clipboard/createPort" {
			typeRequest = 10
		} else if c.Request.Method == "GET" && c.Request.RequestURI == "/clipboard/port" {
			typeRequest = 11
		} else if c.Request.Method == "DELETE" && c.Request.RequestURI == "/clipboard/port" {
			typeRequest = 12
		} else if c.Request.Method == "POST" && c.Request.RequestURI == "/clipboard/port" {
			typeRequest = 13
		} else {
			c.JSON(http.StatusBadRequest, gin.H{})
			c.Abort()
			return
		}

		err := c.BindJSON(&ClipboardJson)
		if err != nil {
			c.IndentedJSON(400, gin.H{"message": "Invalid JSON", "error": err.Error()})
			return
		}

		if antiDDOSRequest(c) {
			c.JSON(http.StatusTooManyRequests, gin.H{})
			c.Abort()
			return
		}
		addLogHttp(typeRequest, c)
		c.Next()
	}
}
