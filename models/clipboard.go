/*
* name : clipboard.go
* package : models
* Defining the clipboard structure
 */

package models

import "github.com/google/uuid"

type Clipboard struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Port         int       `json:"port"`
	Message      string    `json:"message"`
	Password     string    `json:"password"`
	IpOwner      string    `json:"ip_owner"`
	CreationDate string    `json:"creationdate"`
	LastUpdate   string    `json:"lastupdate"`
	LastSee      string    `json:"lastsee"`
}

func (Clipboard) TableName() string {
	return "clipboard"
}
