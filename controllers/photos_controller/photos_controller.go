package photos_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zidanhafiz/go-restapi/helpers/photos_helper"
	"github.com/zidanhafiz/go-restapi/models"
)

func Index(c *gin.Context) {
    photos, err := photos_helper.GetAllPhotos()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Success get all photos","data": photos})
}

func Create(c *gin.Context) {
    var input models.Photo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

		userId := c.MustGet("userId").(string)

    photo, err := photos_helper.CreatePhoto(input, userId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create photo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success create photo","data": photo})
}

func Show(c *gin.Context) {
    photo, err := photos_helper.GetPhotoByID(c.Param("photoId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Photo not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success get photo", "data": photo})
}

func Update(c *gin.Context) {
    var input models.Photo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := photos_helper.UpdatePhotoByID(c.Param("photoId"), input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update photo"})
        return
    }

    photo, err := photos_helper.GetPhotoByID(c.Param("photoId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Photo not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success update photo", "data": photo})
}

func Delete(c *gin.Context) {
    err := photos_helper.DeletePhotoByID(c.Param("photoId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete photo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success delete photo"})
}