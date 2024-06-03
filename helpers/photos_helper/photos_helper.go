package photos_helper

import (
	"strconv"

	"github.com/zidanhafiz/go-restapi/models"
)

func GetAllPhotos() (*[]models.Photo, error) {
    var photos []models.Photo
    result := models.DB.Find(&photos)
    if result.Error != nil {
        return nil, result.Error
    }
    return &photos, nil
}

func GetPhotoByID(id string) (*models.Photo, error) {
    var photo models.Photo
    result := models.DB.Where("id = ?", id).First(&photo)
    if result.Error != nil {
        return nil, result.Error
    }
    return &photo, nil
}

func UpdatePhotoByID(id string, data models.Photo) error {
    var photo models.Photo
    result := models.DB.Model(&photo).Where("id = ?", id).Updates(data).First(&photo)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func CreatePhoto(data models.Photo, userId string) (*models.Photo, error) {
		id, _ := strconv.ParseUint(userId, 10, 0)
		data.UserID = id
    result := models.DB.Create(&data)
    if result.Error != nil{
        return nil, result.Error
    }
    return &data, nil
}

func DeletePhotoByID(photoId string) error {
    var photo models.Photo
    result := models.DB.Where("id = ?", photoId).Delete(&photo)
    if result.Error != nil {
        return result.Error
    }
    return nil
}