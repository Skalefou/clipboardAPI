/*
* name : clipboard.go
* package : models
* Defining the clipboard structure
 */

package models

type Clipboard struct {
	ID           int    `json:"id"`
	Port         int    `json:"port"`
	Message      string `json:"message"`
	Password     string `json:"password"`
	IpOwner      string `json:"ip_owner"`
	CreationDate string `json:"creationdate"`
	LastUpdate   string `json:"lastupdate"`
	LastSee      string `json:"lastsee"`
}

func (Clipboard) TableName() string {
	return "clipboard"
}
