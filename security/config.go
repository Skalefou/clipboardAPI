/*
* name : cleanData.go
* package : security
* Loads environment variables for program operation (configuration and database connection)
 */

package security

import (
	"clipboardAPI/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var DSN string = "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
var ANTI_DDOS_TIME int
var CLIP_DELETION_TIME_WITHOUT_PASSWORD int
var CLIP_DELETION_TIME_WITH_PASSWORD int
var LOG_DELETION_TIME int
var PORT_APP string
var Db *gorm.DB
var ClipboardJson models.Clipboard

var lastIdLog uuid.UUID

// Checks whether the values of environment variables managing time are consistent
func verifyErrorInt(err *error, message string, value int) {
	if err == nil || value < 0 {
		log.Fatal(message)
	}
}

// Load configuration via "config.env" file
func LoadConfig() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("The \"config.env\" file is invalid")
	}

	DSN = fmt.Sprintf(DSN, os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"))

	Db, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failure to open the database, look at the \"config.env\" file")
	}

	ANTI_DDOS_TIME, err = strconv.Atoi(os.Getenv("ANTI_DDOS_TIME"))
	verifyErrorInt(&err, "Anti DDOS time is wrong", ANTI_DDOS_TIME)

	CLIP_DELETION_TIME_WITHOUT_PASSWORD, err = strconv.Atoi(os.Getenv("CLIP_DELETION_TIME_WITHOUT_PASSWORD"))
	verifyErrorInt(&err, "Clip deletion time without password is wrong", CLIP_DELETION_TIME_WITHOUT_PASSWORD)

	CLIP_DELETION_TIME_WITH_PASSWORD, err = strconv.Atoi(os.Getenv("CLIP_DELETION_TIME_WITH_PASSWORD"))
	verifyErrorInt(&err, "Clip deletion time with password is wrong", CLIP_DELETION_TIME_WITH_PASSWORD)

	LOG_DELETION_TIME, err = strconv.Atoi(os.Getenv("LOG_DELETION_TIME"))
	verifyErrorInt(&err, "Clip deletion time with password is wrong", LOG_DELETION_TIME)

	LOG_DELETION_TIME, err = strconv.Atoi(os.Getenv("LOG_DELETION_TIME"))
	verifyErrorInt(&err, "Clip deletion time with password is wrong", LOG_DELETION_TIME)

	PORT_APP = os.Getenv("PORT_APP")
}
