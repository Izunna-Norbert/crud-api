package controllers

import (
	"github.com/Izunna-Norbert/crud-api/initializers"
	"github.com/Izunna-Norbert/crud-api/models"
	"github.com/gin-gonic/gin"
)


func PostCreate(g *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	g.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		g.Status(400)
		return
	}

	g.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(g *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	g.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(g *gin.Context) {

	id := g.Param("id")
	var post models.Post

	result := initializers.DB.Find(&post, id)

	if result.Error != nil {
		g.Status(400)
		return
	}

	g.JSON(200, gin.H{
		"posts": post,
	})
}

func UpdatePost(g *gin.Context) {

	id := g.Param("id")
	var body struct {
		Title string
		Body  string
	}

	g.Bind(&body)

	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{ Title: body.Title, Body: body.Body })

	g.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(g *gin.Context) {
	id := g.Param("id")

	var post models.Post

	initializers.DB.Find(&post, id)

	initializers.DB.Delete(&post, id)

	g.JSON(200, gin.H{
		"post": post,
	})
}