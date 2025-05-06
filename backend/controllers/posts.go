package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"backend/models"

	"github.com/gin-gonic/gin"
)

// GET /posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	userID := c.Query("user_id")
	query := models.DB.Preload("User").Order("created_at desc")
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	query.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

// POST /posts
func CreatePost(c *gin.Context) {
	userID := c.GetInt("user_id")
	content := c.PostForm("content")
	imageURL := ""
	file, err := c.FormFile("image")
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%d%s", userID, time.Now().Unix(), filepath.Ext(file.Filename))
		uploadPath := filepath.Join("uploads", filename)
		os.MkdirAll("uploads", 0755)
		if err := c.SaveUploadedFile(file, uploadPath); err == nil {
			imageURL = "/" + uploadPath
		}
	}
	post := models.Post{
		UserID:    uint(userID),
		Content:   content,
		ImageURL:  imageURL,
		Likes:     0,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	models.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}

// PUT /posts/:id
func UpdatePost(c *gin.Context) {
	userID := c.GetInt("user_id")
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if post.UserID != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not your post"})
		return
	}
	content := c.PostForm("content")
	if content != "" {
		post.Content = content
	}
	removeImage := c.PostForm("remove_image")
	file, err := c.FormFile("image")
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%d%s", userID, time.Now().Unix(), filepath.Ext(file.Filename))
		uploadPath := filepath.Join("uploads", filename)
		os.MkdirAll("uploads", 0755)
		if err := c.SaveUploadedFile(file, uploadPath); err == nil {
			// Delete old image if exists
			if post.ImageURL != "" {
				os.Remove(post.ImageURL[1:]) // remove leading slash
			}
			post.ImageURL = "/" + uploadPath
		}
	} else if removeImage == "true" {
		if post.ImageURL != "" {
			os.Remove(post.ImageURL[1:]) // remove leading slash
		}
		post.ImageURL = ""
	}
	post.UpdatedAt = time.Now().Unix()
	models.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

// DELETE /posts/:id
func DeletePost(c *gin.Context) {
	userID := c.GetInt("user_id")
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if post.UserID != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not your post"})
		return
	}
	// Delete image file if exists
	if post.ImageURL != "" {
		os.Remove(post.ImageURL[1:]) // remove leading slash
	}
	models.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

// POST /posts/:id/like
func LikePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	post.Likes++
	models.DB.Save(&post)
	c.JSON(http.StatusOK, gin.H{"likes": post.Likes})
}
