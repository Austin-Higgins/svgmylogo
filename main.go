package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	ID   uint
	Face string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Controllers
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/animals", func(c *gin.Context) {
		var animals []Animal
		result := db.Find(&animals)
		c.HTML(http.StatusOK, "index.go.tmpl", gin.H{
			"Animals": result,
		})
	})

	router.GET("/animals/:id", func(c *gin.Context) {
		id := c.Param("id")
		var animal Animal
		db.First(&animal, id)

		c.HTML(http.StatusOK, "show.go.tmpl", gin.H{
			"Face": animal.Face,
		})
	})

	router.Run(":8080")
}
