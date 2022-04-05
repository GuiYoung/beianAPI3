package v1

import (
	"beianAPI/model"
	"beianAPI/utils/crawler"
	"beianAPI/utils/errmsg"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// get info (token)
func GetInfo(c *gin.Context) {
	domainName := c.Query("domainName")
	beian, err := model.FindBeianByName(domainName)
	code := errmsg.SUCCESS

	if errors.Is(err, gorm.ErrRecordNotFound) {
		beian = crawler.GetBeiAnInfo(domainName)
		if beian == nil {
			code = errmsg.ERROR_GET_INFO_FAILED
		} else {
			_ = model.InsertBeian(beian)
		}
	} else if err != nil {
		code = errmsg.ERROR_GET_INFO_FAILED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"info":    beian,
	})
}

// get latest info (token)
func GetLatestInfo(c *gin.Context) {
	domainName := c.Query("domainName")
	var beian *model.Beian
	code := errmsg.SUCCESS

	beian = crawler.GetBeiAnInfo(domainName)
	if beian == nil {
		code = errmsg.ERROR_GET_INFO_FAILED
	} else {
		_ = model.InsertBeian(beian)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"info":    beian,
	})
}
