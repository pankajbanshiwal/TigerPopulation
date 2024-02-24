package Users

import (
	"TigerPopulation/Domain/Users"
	services "TigerPopulation/Services/Users"
	"TigerPopulation/Utils"
	"net/http"
	"net/mail"
	"regexp"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user Users.CreateUser
	var resp Utils.ApiResponse
	var err error
	var result *Users.CreateUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		resp.Status = false
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	result, err = services.Login(&user)
	if err != nil {
		resp.Status = false
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	ctx.JSON(http.StatusOK, result)

}

func CreateUser(c *gin.Context) {
	var resp Utils.ApiResponse
	var request Users.CreateUser
	if err := c.ShouldBindJSON(&request); err != nil {
		resp.Status = false
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// check if user name contains any whitespaces
	if regexp.MustCompile(`\s`).MatchString(request.UserName) {
		resp.Status = false
		resp.Code = http.StatusNotAcceptable
		resp.Message = "Whitespaces are not allowed in username"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}

	// Validate password
	err := Utils.VerifyPassword(request.Password)
	if err != nil {
		resp.Status = false
		resp.Code = http.StatusNotAcceptable
		resp.Message = err.Error()
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}

	// validate email id
	_, err = mail.ParseAddress(request.Email)
	if err != nil {
		resp.Status = false
		resp.Code = http.StatusNotAcceptable
		resp.Message = err.Error()
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}

	err = services.CreateUser(&request)
	if err != nil {
		resp.Status = false
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = true
	resp.Code = http.StatusOK
	resp.Message = "Success"
	c.JSON(http.StatusOK, resp)
}
