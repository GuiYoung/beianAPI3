package v1

import (
	"beianAPI/model"
	"beianAPI/utils/errmsg"
	"beianAPI/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

// sign up (username password)
func SignUp(c *gin.Context) {

	var user model.User
	_ = c.ShouldBindJSON(&user)
	msg, err := validator.Validate(&user)
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  errmsg.ERROR,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	code := model.CreateUser(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// get token (username password)
func GetToken(c *gin.Context) {
	userName := c.PostFormArray("userName")
	pwd := c.PostFormArray("password")

	token, code := model.CheckUser(userName[0], pwd[0])

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
