/*
* name : cleanData.go
* package : security
* Cleans up clipboards that have been left unused for far too long (7 days for clips without a password, and
* have no password, and 14 days for those that do).
 */

package security

import (
	"clipboardAPI/models"
	"time"
)

func CleanData() {

	if CLIP_DELETION_TIME_WITHOUT_PASSWORD > 0 {
		Db.Where("last_see < ? AND password = ?", time.Now().AddDate(0, 0, -CLIP_DELETION_TIME_WITHOUT_PASSWORD), "").Delete(&models.Clipboard{})
	}
	if CLIP_DELETION_TIME_WITH_PASSWORD > 0 {
		Db.Where("last_see < ?", time.Now().AddDate(0, 0, -CLIP_DELETION_TIME_WITH_PASSWORD)).Delete(&models.Clipboard{})
	}
	if LOG_DELETION_TIME > 0 {
		Db.Where("date_request < ?", time.Now().AddDate(0, 0, -LOG_DELETION_TIME)).Delete(&models.Log{})
	}

	log := models.Log{
		TypeRequest: 1,
		Clipboard:   ClipboardJson.ID,
		DateRequest: time.Now().Format("2006-01-02 15:04:05"),
	}
	Db.Create(&log)
}
