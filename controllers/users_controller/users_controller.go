package users_controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zidanhafiz/go-restapi/helpers/users_helper"
	"github.com/zidanhafiz/go-restapi/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := users_helper.GetUserByUsernameOrEmail(input.Username, input.Email)
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	user, err := users_helper.CreateUser(input.Username, input.Email, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success register user", "data": user})
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := users_helper.GetUserByUsernameOrEmail(input.Username, input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   strconv.FormatUint(uint64(user.ID), 10),
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("pbi-rakamin"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session_token",
		Value:    tokenString,
		HttpOnly: true,
		Path: 	 "/",
	})

	c.JSON(http.StatusOK, gin.H{"message": "success login user"})
}

func Index(c *gin.Context) {
	users, err := users_helper.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success get all users", "data": users})
}

func Show(c *gin.Context) {
	user, err := users_helper.GetUserByID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success get user", "data": user})
}

func Update(c *gin.Context) {
	_, err := users_helper.GetUserByID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := users_helper.GetUserByUsernameOrEmail(input.Username, input.Email); err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Email already exists"})
		return
	}

	if err := users_helper.UpdateUserByID(c.Param("userId"), input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success update user"})
}

func Delete(c *gin.Context) {
	_, err := users_helper.GetUserByID(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := users_helper.DeleteUserByID(c.Param("userId")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success delete user"})
}
