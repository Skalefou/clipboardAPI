/*
* name : log.go
* package : models
* Defining the log structure
 */

package models

import "github.com/google/uuid"

type Log struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	TypeRequest int       `json:"typerequest"`
	IpUser      string    `json:"ipuser"`
	Clipboard   uuid.UUID `gorm:"type:uuid" json:"clipboard"`
	DateRequest string    `json:"daterequest"`
	Active      bool      `json:"active"`
}

func (Log) TableName() string {
	return "log"
}
