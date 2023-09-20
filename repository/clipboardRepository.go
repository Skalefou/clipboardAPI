/*
* name : clipboardRepository.go
* package : models
* Manages the various interactions between the Clipboard structure and the database
 */

package repository

import (
	"clipboardAPI/models"
	"clipboardAPI/security"
	"math/rand"
	"time"
)

// Determines a port by randomly assigning an unused port of up to 6 digits.
func GenPort() int {
	var portGenerate int
	var portUse []int
	var unused [999999]bool

	security.Db.Table("clipboard").Select("port").Scan(&portUse)

	for _, num := range portUse {
		unused[num] = true
	}

	for {
		portGenerate = 1 + rand.Intn(999999)
		if !unused[portGenerate] {
			break
		}
	}

	return portGenerate
}

// Determines if the ip is already in use
func NoDuplicateIP(ip string, clip *models.Clipboard) bool {
	security.Db.Where("ip_owner = ?", ip).First(&clip)
	return clip.ID == 0
}

// Create a new clipboard record in the database
func CreateClip(ip string, port int, password string) models.Clipboard {
	clipboard := models.Clipboard{
		ID:           port,
		Message:      "",
		Password:     password,
		IpOwner:      ip,
		CreationDate: time.Now().Format("2006-01-02 15:04:05"),
		LastUpdate:   time.Now().Format("2006-01-02 15:04:05"),
		LastSee:      time.Now().Format("2006-01-02 15:04:05"),
	}
	security.Db.Create(&clipboard)
	security.ActiveRequestLog()
	clipboard.Password = ""
	return clipboard
}

// Retrieves a clipboard according to its id and updates the "last_see" column
func GetClipboard(clipboard *models.Clipboard) models.Clipboard {
	var clipboardTemp models.Clipboard
	security.Db.First(&clipboardTemp, clipboard.ID)
	security.Db.Model(&models.Clipboard{}).Where("id = ?", clipboardTemp.ID).Update("last_see", time.Now().Format("2006-01-02 15:04:05"))
	security.ActiveRequestLog()
	return clipboardTemp
}

// Delete a clipboard using a password
func DeleteClipboardFromPassword(clipboard *models.Clipboard) bool {
	var password string
	if security.Db.Model(&models.Clipboard{}).Where("id = ?", clipboard.ID).Select("password").Scan(&password).RowsAffected == 0 {
		return false
	}
	if password == clipboard.Password && password != "" {
		security.ActiveRequestLog()
		security.Db.Where("id", clipboard.ID).Delete(&models.Clipboard{})
		return true
	} else {
		return false
	}
}

// Updates the database, in particular the "Message" column
func UpdateClipboard(clipboard models.Clipboard) bool {
	var password string
	if security.Db.Model(&models.Clipboard{}).Where("id = ?", clipboard.ID).Select("password").Scan(&password).RowsAffected == 0 {
		return false
	}

	if clipboard.Password == password {
		security.ActiveRequestLog()
		security.Db.Model(&clipboard).UpdateColumns(models.Clipboard{
			Message:    clipboard.Message,
			LastUpdate: time.Now().Format("2006-01-02 15:04:05"),
			LastSee:    time.Now().Format("2006-01-02 15:04:05"),
		})
		return true
	}
	return false
}
