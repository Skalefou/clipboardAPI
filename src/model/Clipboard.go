package model

import "time"

type Clipboard struct {
	UUID           string    `json:"uuid"`
	IPAuthor       string    `json:"ip_author"`
	IDAccess       int       `json:"id_access"`
	Password       string    `json:"password"`
	DateCreation   time.Time `json:"date_creation"`
	DateLastUpdate time.Time `json:"date_last_update"`
	Message        string    `json:"message"`
}
