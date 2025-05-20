package main

import (
	"fmt"
	"log"
	"os"
	"todolist/middleware"
	ginitem "todolist/module/item/transport/gin"
	"todolist/module/upload"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Connect DB failed:", err)
		return
	}
	fmt.Println("Currently connected to DB")

	r := gin.Default()
	r.Static("/static", "./static")
	r.Use(middleware.Recover())

	v1 := r.Group("/v1")
	{
		v1.PUT("/upload", upload.Upload(db))

		items := v1.Group("/items")
		{
			items.GET("/", ginitem.GetItems(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.POST("/", ginitem.CreateItem(db))
			items.PUT("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	r.Run(os.Getenv("APP_HOST")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
