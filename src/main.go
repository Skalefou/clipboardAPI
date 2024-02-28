package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	_ = router.Run("localhost:8080")
}
