package controller

import (
	"github.com/gin-gonic/gin"
	// "github.com/sayalipol/go-crud/initializers"
	"github.com/sayalipol/go-crud/initializers"
	"github.com/sayalipol/go-crud/models"
)

// Create
func PostControllers(c *gin.Context) {

	// create post by hardcoding data WAY 1:-
	// post := models.Post{Title: "Blog 1", Body: "I am Body of first blog"}
	// result := initializers.DB.Create(&post)
	// if result.Error != nil {
	// 	c.Status(400)
	// 	return
	// }
	// // return it
	// c.JSON(200, gin.H{
	// 	"post": post,
	// })

	// WAY 2:-   taking data off request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// Read
// Read all posts
func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// Read single post
func PostsShow(c *gin.Context) {
	//get id of the post
	id := c.Param("id")

	//Get the post
	var post []models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// get id from URL
	id := c.Param("id")

	//get data off rquest body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//UPDATE IT
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	///Get ID to be delted
	id := c.Param("id")

	// delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.Status(200)
}
