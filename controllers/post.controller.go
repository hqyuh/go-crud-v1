package controllers

import (
	"github.com/gin-gonic/gin"
	"go-crud-v1/service"
)

type PostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func PostCreate(c *gin.Context) {
	service.PostCreate(c)
}

func FindAll(c *gin.Context) {
	service.FindAll(c)
}

func FindById(c *gin.Context) {
	service.FindById(c)
}

func PostUpdate(c *gin.Context) {
	service.PostUpdate(c)
}

func DeleteById(c *gin.Context) {
	service.DeleteById(c)
}
