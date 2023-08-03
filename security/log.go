/*
* name : log.go
* package : security
* Manages logs
 */

package security

import (
	"clipboardAPI/models"
	"github.com/gin-gonic/gin"
	"time"
)

// Adds a log record to the database
func addLogHttp(typeRequestInput int, c *gin.Context) {
	log := models.Log{
		TypeRequest: typeRequestInput,
		IpUser:      c.Request.RemoteAddr,
		Clipboard:   ClipboardJson.ID,
		DateRequest: time.Now().Format("2006-01-02 15:04:05"),
	}

	Db.Create(&log)
	lastIdLog = log.ID
}

// Activate query in logs
func ActiveRequestLog() {
	Db.Model(&models.Log{}).Where("id = ?", lastIdLog).Update("active", true)
}
