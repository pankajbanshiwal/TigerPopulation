package accessToken

import (
	"TigerPopulation/Utils/dbConfig"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golang/glog"
)

// ParseToken parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (userid float64, Error error) {
	dbConfig := dbConfig.ViperConfigDev()
	SecretKey := []byte(dbConfig.JwtSecretKey)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userid := claims["userid"].(float64)
		var tm time.Time
		switch iat := claims["exp"].(type) {
		case float64:
			tm = time.Unix(int64(iat), 0)
		case json.Number:
			v, _ := iat.Int64()
			tm = time.Unix(v, 0)
		}
		glog.V(3).Infoln("time==", tm)

		return userid, nil
	} else {
		return 0, err
	}
}

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(userid int) (string, error) {
	dbConfig := dbConfig.ViperConfigDev()
	SecretKey := []byte(dbConfig.JwtSecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["userid"] = userid
	claims["exp"] = time.Now().AddDate(1, 0, 0).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		glog.V(3).Infoln("Error in Generating key")
		return "", err
	}

	return tokenString, nil
}
