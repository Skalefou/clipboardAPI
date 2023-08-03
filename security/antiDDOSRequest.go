/*
* name : antiDDOSRequest.go
* package : security
* Manages anti-DDOS by determining whether too many requests are coming from an IP address
 */

package security

import (
	"clipboardAPI/models"
	"github.com/gin-gonic/gin"
	"time"
)

func antiDDOSRequest(c *gin.Context) bool {
	var b int64
	Db.Model(&models.Log{}).Where("date_request > ? AND ip_user = ?", time.Now().Add(time.Duration(-ANTI_DDOS_TIME)*time.Millisecond), c.Request.RemoteAddr).Count(&b)
	return b > 0
}
