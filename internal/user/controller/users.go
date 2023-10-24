package usercontroller

import (
	"gin-gorm-rest-api/config"
	"gin-gorm-rest-api/internal/user/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func GetUser(c *gin.Context) {
	users := []models.Users{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {
	var user models.Users
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed read body",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash to password",
		})
		return
	}
	user.Password = string(hash)
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var body models.Users

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed read body",
		})
		return
	}

	var user models.Users
	config.DB.First(&user, "email = ?", body.Email)

	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Not Found Users",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email and password",
		})
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("huongdao"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 24*30*60, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.Users
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.Users
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func CheckLoggin(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
