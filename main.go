package main

import (
	"clipboardAPI/controller"
	"clipboardAPI/security"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func main() {
	security.LoadConfig()
	router := gin.Default()

	c := cron.New()
	err := c.AddFunc("0 0 3 * * *", security.CleanData)
	if err != nil {
		return
	}
	c.Start()

	controller.HttpSource(router)

	router.Run(security.PORT_APP)
}
