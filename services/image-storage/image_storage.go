package image_storage

import (
	"apricotka.com.ua/packages/models"
	"apricotka.com.ua/packages/models/entities"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type productImageFormData struct {
	ProductId uint64                `form:"product_id" binding:"required"`
	File      *multipart.FileHeader `form:"file" binding:"required"`
}

type ImgurResponse struct {
	Data struct {
		Link string `json:"link"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

func GetProductImages(c *gin.Context) {
	var productImages []entities.ProductImage
	result := models.DB.Find(&productImages)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, result.Error.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, productImages)
}

func UploadImage(c *gin.Context) {
	var formData productImageFormData
	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := uploadToImgur(c, formData.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productId := formData.ProductId

	productImage := entities.ProductImage{
		ProductId: productId,
		File:      link,
	}

	result := models.DB.Create(&productImage)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, productImage)
}

func UpdateImage(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")

	// Bind the response body
	var formData productImageFormData
	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the product image with the specified ID
	var productImage entities.ProductImage
	if err := models.DB.First(&productImage, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Upload file to imgur
	link, err := uploadToImgur(c, formData.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the product image in the database
	productImage.ProductId = formData.ProductId
	productImage.File = link
	if err := models.DB.Save(&productImage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, productImage)
}

func uploadToImgur(c *gin.Context, file *multipart.FileHeader) (string, error) {
	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}(src)

	// Create a new multipart writer and write the file
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("image", file.Filename)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, src)
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://api.imgur.com/3/image", &body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", os.Getenv("IMGUR_CLIENT_ID")))
	req.Header.Set("Token", fmt.Sprintf("Bearer %s", os.Getenv("IMGUR_TOKEN")))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}(resp.Body)

	// Read the response
	var imgurResponse ImgurResponse
	err = json.NewDecoder(resp.Body).Decode(&imgurResponse)
	if err != nil {
		return "", err
	}

	// Check for errors in the response
	if !imgurResponse.Success {
		return "", fmt.Errorf("imgur upload failed with status code %d", imgurResponse.Status)
	}

	// Return the link to the uploaded image
	return imgurResponse.Data.Link, nil
}

func DeleteImage(c *gin.Context) {
	// Get the ID from the URL parameter
	id := c.Param("id")

	// Find the product image with the specified ID
	var productImage entities.ProductImage
	if err := models.DB.First(&productImage, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Delete the product image from the database
	if err := models.DB.Delete(&productImage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product image deleted successfully"})
}
