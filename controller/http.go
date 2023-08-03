/*
* name : http.go
* package : controller
* Also controller, manages various routes
 */

package controller

import (
	"clipboardAPI/security"
	"clipboardAPI/service"
	"github.com/gin-gonic/gin"
)

func HttpSource(router *gin.Engine) {
	clipboard := router.Group("/clipboard")

	clipboard.Use(security.MiddlewareClipboard())

	clipboard.POST("/createPort", service.CreatePort)
	clipboard.GET("/port", service.GetPort)
	clipboard.DELETE("/port", service.DeletePort)
	clipboard.POST("/port", service.UpdatePort)
}
