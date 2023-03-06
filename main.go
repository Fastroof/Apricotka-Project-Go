package main

import (
	"apricotka.com.ua/packages/models"
	"apricotka.com.ua/packages/services/gallery"
	"apricotka.com.ua/packages/services/messenger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load env variables from file
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Gin router
	r := gin.Default()

	// Initialize the database
	models.ConnectDatabase()

	// Set up the routes for each module
	r.GET("/api/gallery/gallery", gallery.Gallery)
	r.GET("/api/gallery/gallery-images", gallery.Images)
	r.POST("/api/gallery/new/gallery", gallery.NewGallery)
	r.POST("/api/gallery/new/image", gallery.NewGalleryImage)
	r.PUT("/api/gallery/gallery/:id", gallery.UpdateGallery)
	r.PUT("/api/gallery/image/:id", gallery.UpdateGalleryImage)
	r.DELETE("/api/gallery/gallery/:id", gallery.DeleteGallery)
	r.DELETE("/api/gallery/image/:id", gallery.DeleteGalleryImage)

	r.GET("/api/messenger/ping", messenger.Ping)
	r.POST("/api/messenger/send/mail", messenger.SendMail)

	//r.GET("/orders", orders.Handler)
	//
	//r.GET("/security/login", security.HandleLogin)
	//r.POST("/security/login", security.HandleLogin)
	//r.GET("/security/logout", security.HandleLogout)
	//
	//r.GET("/storage", storage.Handler)

	// Start the server
	log.Println("Starting server on port 8080")
	err = r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
