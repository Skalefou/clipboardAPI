// Program Launch Package
package main

import (
	"github.com/gin-gonic/gin"
)

// Program Launch function
func main() {
	router := gin.Default()

	_ = router.Run("localhost:8080")
}
