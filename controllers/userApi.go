package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harish1907/zgolang/initializers"
	"github.com/harish1907/zgolang/models"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath

// PingExample godoc
// @Summary ping example
// @Schemes
// @Tags example
// @Accept json
// @Produce json
// @Param        body  body      models.UserModel  true  "UserModel JSON"
// @Success 200 {json} SignUp
// @Router /singup/user/ [post]
func SignUp(c *gin.Context) {
	// Get the email and password in request body.
	var body struct {
		Email string `json:"email"`
		Password string  `json:"password"`
	}

	c.Bind(&body)

	if body.Email == "" || body.Password == ""{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Email and password is required field",
			"status_code": 404,
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message": err,
			"status_code": 404,
		})
		return
	}

	user := &models.UserModel{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "This user already exist.",
			"status_code": 404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User creating successfully.",
		"status_code": 200,
	})
}


// @BasePath 

// PingExample godoc
// @Summary ping example
// @Schemes
// @Tags example
// @Accept json
// @Produce json
// @Param        body  body      models.UserModel  true  "UserModel JSON"
// @Success 200 {json} LoginAPI
// @Router /login/user/ [post]
func LoginAPI(c *gin.Context){
	var body struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	c.Bind(&body)

	if body.Email == "" || body.Password == ""{
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Email and password is required field",
			"status_code": 404,
		})
		return
	}

	var user models.UserModel
	initializers.DB.First(&user, "email=?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid email.",
			"status_code": 404,
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Incorrect password..",
			"status_code": 404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User login successfully.",
		"status_code": 200,
	})
}