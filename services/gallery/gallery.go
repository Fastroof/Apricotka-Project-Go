package gallery

import (
	"apricotka.com.ua/packages/models"
	"apricotka.com.ua/packages/models/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type galleryFormData struct {
	Title string `form:"title" binding:"required"`
}

type galleryImageFormData struct {
	GalleryId uint64 `form:"gallery_id" binding:"required"`
	File      string `form:"file"       binding:"required"`
}

func Gallery(c *gin.Context) {
	var galleries []entities.Gallery
	// Get all galleries
	result := models.DB.Find(&galleries)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, galleries)
}

func Images(c *gin.Context) {
	var images []entities.GalleryImage
	// Get all images
	result := models.DB.Find(&images)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, images)
}

func NewGallery(c *gin.Context) {
	var formData galleryFormData
	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := formData.Title
	if len(title) != 0 {
		gallery := entities.Gallery{
			Title: title,
		}
		result := models.DB.Create(&gallery)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gallery)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Length of title must be > 0"})
}

func NewGalleryImage(c *gin.Context) {
	var formData galleryImageFormData
	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file := formData.File
	galleryId := formData.GalleryId

	if len(file) != 0 {
		var exists bool
		err := models.DB.Model(entities.Gallery{}).
			Select("count(*) > 0").
			Where("id = ?", galleryId).
			Find(&exists).
			Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		galleryImage := entities.GalleryImage{
			GalleryId: galleryId,
			File:      file,
		}

		result := models.DB.Create(&galleryImage)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, galleryImage)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Length of file must be > 0"})
}

// UpdateGallery updates a gallery by ID
func UpdateGallery(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")

	// Bind the request body
	var formData galleryFormData
	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the gallery with the specified ID
	var gallery entities.Gallery
	if err := models.DB.First(&gallery, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Update the gallery in the database
	gallery.Title = formData.Title
	if err := models.DB.Save(&gallery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gallery)
}

func UpdateGalleryImage(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")

	// Bind the request body
	var formData galleryImageFormData
	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the gallery image with the specified ID
	var image entities.GalleryImage
	if err := models.DB.First(&image, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Update the gallery image in the database
	image.GalleryId = formData.GalleryId
	image.File = formData.File
	if err := models.DB.Save(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, image)
}

func DeleteGallery(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")

	// Find the gallery with the specified ID
	var gallery entities.Gallery
	if err := models.DB.First(&gallery, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Delete the gallery from the database
	if err := models.DB.Delete(&gallery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gallery deleted successfully"})
}

func DeleteGalleryImage(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")

	// Find the gallery with the specified ID
	var image entities.GalleryImage
	if err := models.DB.First(&image, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Delete the gallery from the database
	if err := models.DB.Delete(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gallery image deleted successfully"})
}
