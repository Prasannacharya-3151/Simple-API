package handlers

import (
	"net/http"
	"simple-api/config"
	"simple-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/utils"
)

//signup
func Signup(c *gin.Context) {  //c *gin.Context which containes a request, header, body, response
	var input models.SignupInput //which creates a empty struct like a name, email, password in the empty string
    //we can wrie like this also 
	//if err :=
	//var err error
	//err = someFunction()
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if user alredy exists
	var  existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registerd"})
		return
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Passowrd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User {
		Name: input.Name,
		Email: input.Email,
		Password: string(hashedPassword),
	}
    //create user
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"token": token,
		"user": gin.H{
			"id": user.ID,
			"name": user.Name,
			"email": user.Email,
		},
	})
}

func Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email=?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	//compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Login successful",
		"token": token,
		"user": gin.H{
			"id": user.ID,
			"name": user.Name,
			"email": user.Email,
		},
	})
}

//profile (protected route example)
func Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")

	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to your profile",
		"user_id": userID,
		"email":email,
	})
}