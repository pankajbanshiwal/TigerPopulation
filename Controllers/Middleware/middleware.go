package Middleware

import (
	"TigerPopulation/Utils/accessToken"
	"context"
	"strings"

	//"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type CustomResult struct {
	Status        bool      `json:"status"`
	Code          int       `json:"code"`
	Message       string    `json:"message"`
	Expirytime    time.Time `json:"expirytime"`
	Sessionstatus bool      `json:"sessionstatus"`
}

func CustomTokenMiddleware(contextKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.FullPath(), "user/create") {
			c.Next()
			return
		}
		if strings.Contains(c.FullPath(), "user/login") {
			c.Next()
			return
		}
		token := c.Request.Header.Get("token")
		os.Setenv("token", token)
		glog.V(3).Infoln("tkn", os.Getenv("token"))
		if token == "" {
			var res CustomResult
			res.Code = http.StatusUnauthorized
			res.Status = false
			res.Message = "Empty token"
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		id, err := accessToken.ParseToken(token)
		if err != nil {
			var res CustomResult
			res.Code = 401
			res.Status = false
			res.Message = "Token expired"
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		userid := int(id)
		glog.V(3).Infoln("user_id = ", userid)
		if userid == 0 {
			var res CustomResult
			res.Code = 401
			res.Status = false
			res.Message = "User not found"
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}
		ctx := context.WithValue(c.Request.Context(), contextKey, userid)

		c.Request = c.Request.WithContext(ctx)

		c.Next()

	}
}

// func CustomSessionValidator() (*users.User, error) {

// 	var user *users.User
// 	var err error

// 	var userId int

// 	token := os.Getenv("token")

// 	if token == "" {
// 		return user, errors.New("token empty")
// 	}

// 	id, err1 := accessToken.ParseToken(token)

// 	if err1 != nil {
// 		return user, errors.New("token expired")
// 	}
// 	userId = int(id)

//		user = users.GetSignUpResponse(userId)
//		if user.Userid == 0 {
//			return user, errors.New("no user found")
//		}
//		return user, err
//	}
func SetSessionToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		os.Setenv("token", token)
		glog.V(3).Infoln("tkn", os.Getenv("token"))
		if token == "" {
			c.Next()
			return
		}

	}
}
