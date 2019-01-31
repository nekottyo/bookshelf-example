package main

import (
	"reflect"
	"strconv"

	"github.com/nekottyo/bookshelf/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}

		if id <= 0 {
			c.JSON(400, gin.H{"status": "id should be bigger than 0"})
			return
		}

		ctrl := controllers.NewUser()

		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, result)
	})

	r.Run(":8080")
}
