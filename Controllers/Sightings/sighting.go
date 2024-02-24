package Sightings

import (
	"TigerPopulation/Domain/Sightings"
	"TigerPopulation/Domain/Tigers"
	services "TigerPopulation/Services/Sightings"
	"TigerPopulation/Utils"
	"TigerPopulation/Utils/accessToken"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSighting(c *gin.Context) {
	var resp Utils.ApiResponse
	var request Sightings.Sighting
	if err := c.ShouldBindJSON(&request); err != nil {
		resp.Status = false
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	token := c.Request.Header.Get("token")
	userId, err := accessToken.ParseToken(token)
	if err != nil {
		resp.Status = false
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if userId == 0 || len(token) == 0 {
		resp.Status = false
		resp.Code = http.StatusUnauthorized
		resp.Message = "Unauthorized"
		c.JSON(http.StatusUnauthorized, resp)
		return
	}
	request.UserId = int(userId)

	// check if dob is set
	if request.TigerId == 0 {
		resp.Status = false
		resp.Code = http.StatusNotAcceptable
		resp.Message = "Tiger field is required"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}

	// check if dob is set
	if request.Lat == 0 || request.Long == 0 {
		resp.Status = false
		resp.Code = http.StatusNotAcceptable
		resp.Message = "Last seen location can not be empty"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}

	err = services.CreateSighting(&request)
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

func GetTigerSightings(ctx *gin.Context) {
	var resp Tigers.StdResponse

	token := ctx.Request.Header.Get("token")
	userId, err := accessToken.ParseToken(token)
	if err != nil {
		resp.Status = false
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if userId == 0 || len(token) == 0 {
		resp.Status = false
		resp.Code = http.StatusUnauthorized
		resp.Message = "Unauthorized"
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	// parse tiger id
	tigerId, _ := strconv.Atoi(ctx.Query("tigerId"))
	if tigerId == 0 {
		resp.Status = false
		resp.Code = http.StatusNotAcceptable
		resp.Message = "Tiger id can not be empty"
		ctx.JSON(http.StatusNotAcceptable, resp)
		return
	}

	// parse page number
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page == 0 {
		page = 1
	}

	data, err := services.GetTigerSightings(tigerId, page)
	if err != nil {
		resp.Status = false
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = true
	resp.Code = http.StatusOK
	resp.Message = "Success"
	resp.Data = data
	ctx.JSON(http.StatusOK, resp)
}
