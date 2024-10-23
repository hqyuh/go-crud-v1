package service

import (
	"go-crud-v1/initalizers"
	"go-crud-v1/models"

	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func PostCreate(c *gin.Context) {
	var request PostRequest
	err := c.Bind(&request)

	if err != nil {
		return
	}

	post := models.Post{
		Title: request.Title,
		Body:  request.Body,
	}
	result := initalizers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func FindAll(c *gin.Context) {
	var posts []models.Post
	initalizers.DB.Find(&posts)

	c.JSON(200, gin.H{"posts": posts})
}

func FindById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initalizers.DB.First(&post, id)

	c.JSON(200, gin.H{"post": post})
}

func PostUpdate(c *gin.Context) {
	id := c.Query("id")

	var request PostRequest
	err := c.Bind(&request)

	if err != nil {
		return
	}

	// find by id
	var post models.Post
	initalizers.DB.First(&post, id)
	// update
	initalizers.DB.Model(&post).Updates(models.Post{
		Title: request.Title,
		Body:  request.Body,
	})

	c.JSON(200, gin.H{"post": post})
}

func DeleteById(c *gin.Context) {
	id := c.Param("id")

	initalizers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
