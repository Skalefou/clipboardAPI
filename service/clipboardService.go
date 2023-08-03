/*
* name : clipboardService.go
* package : service
* Clipboard structure service, performs processing on service values
 */

package service

import (
	"clipboardAPI/models"
	"clipboardAPI/repository"
	"clipboardAPI/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create a clipboard, check if a user already has one, check if the hash of the password is valid
func CreatePort(c *gin.Context) {
	ip := c.ClientIP()
	var clip models.Clipboard

	if !(security.ClipboardJson.Password == "" || (len(security.ClipboardJson.Password) >= 60 && len(security.ClipboardJson.Password) <= 72)) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid password"})
		return
	}

	if repository.NoDuplicateIP(ip, &clip) {
		c.JSON(http.StatusOK, gin.H{"clipboard": repository.CreateClip(ip, repository.GenPort(), security.ClipboardJson.Password)})
	} else {
		clip.Password = ""
		c.JSON(http.StatusOK, gin.H{"clipboard": clip})
	}
}

// Recovers clipboard data, censors password and ip
func GetPort(c *gin.Context) {

	password := security.ClipboardJson.Password
	security.ClipboardJson = repository.GetClipboard(&security.ClipboardJson)
	if security.ClipboardJson.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wrong clipboard or password"})
	} else {
		if password != security.ClipboardJson.Password {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wrong clipboard or password"})
			return
		}
		if security.ClipboardJson.IpOwner != c.ClientIP() {
			security.ClipboardJson.IpOwner = ""
		}
		security.ClipboardJson.Password = ""
		c.JSON(http.StatusOK, gin.H{"clipboard": security.ClipboardJson})
	}
}

// Deletes a clipboard
func DeletePort(c *gin.Context) {
	if repository.DeleteClipboardFromPassword(&security.ClipboardJson) {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "wrong clipboard or password or deletion is impossible because clipboard has no password"})
	}
}

// Updates clipboard values
func UpdatePort(c *gin.Context) {
	if repository.UpdateClipboard(security.ClipboardJson) {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusForbidden, gin.H{"message": "wrong clipboard or password"})
}
