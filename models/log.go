/*
* name : log.go
* package : models
* Defining the log structure
 */

package models

type Log struct {
	ID          int    `json:"id"`
	TypeRequest int    `json:"typerequest"`
	IpUser      string `json:"ipuser"`
	Clipboard   int    `json:"clipboard"`
	DateRequest string `json:"daterequest"`
	Active      bool   `json:"active"`
}

func (Log) TableName() string {
	return "log"
}
