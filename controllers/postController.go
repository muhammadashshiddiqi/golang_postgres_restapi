package controllers

import (
	"go-crud-postgresql/initializers"
	model "go-crud-postgresql/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//get data req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//create a post
	post := model.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"message": "data has been created",
		"data":    post,
	})
}

func PostsAllData(c *gin.Context) {
	//get the posts
	var lsPost []model.Post
	result := initializers.DB.Find(&lsPost)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"data": lsPost,
	})
}

func PostsShow(c *gin.Context) {
	//get id url
	id := c.Param("id")

	//get the posts
	var post model.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//get id url
	id := c.Param("id")

	//get data req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//find post updating
	var post model.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(model.Post{Title: body.Title, Body: body.Body})

	//response with it
	c.JSON(200, gin.H{
		"message": "data has been updated",
		"data":    post,
	})
}

func PostsDelete(c *gin.Context) {
	//get id url
	id := c.Param("id")

	//find the post
	initializers.DB.Delete(&model.Post{}, &id)

	//response
	c.JSON(200, gin.H{
		"message": "data has been deleted",
	})

}
