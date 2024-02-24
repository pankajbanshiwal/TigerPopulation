package Utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"
	"time"
	"unicode"
)

type ApiResponse struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const EarthRadius = 6371 // in kilometers

func GetDateTime() string {
	current_time := time.Now()
	datetime := current_time.Format("2006-01-02 15:04:05")
	return datetime
}

func CalculateDistanceBetweenLocations(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0
	a := math.Pow(math.Sin(dLat/2.0), 2.0) + math.Cos(lat1*math.Pi/180.0)*math.Cos(lat2*math.Pi/180.0)*math.Pow(math.Sin(dLon/2.0), 2.0)
	c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))
	return EarthRadius * c
}

func VerifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("uppercase letter missing")
	}
	if !numberPresent {
		appendError("atleast one numeric character required")
	}
	if !specialCharPresent {
		appendError("special character missing")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}

func UploadFileToDigitalOcean() {
	// Create a new http.Client
	client := &http.Client{}

	var fileContents []byte // this should contain file content
	// Create a new http.Request
	req, err := http.NewRequest("POST", "https://[YOUR_DIGITALOCEAN_SPACES_ENDPOINT]/[YOUR_BUCKET_NAME]/[IMAGE_NAME]", bytes.NewReader(fileContents))
	if err != nil {
		log.Fatal(err)
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "image/jpeg")

	// Add your DigitalOcean Spaces access token to the Authorization header
	req.Header.Set("Authorization", "Bearer [YOUR_DIGITALOCEAN_SPACES_ACCESS_TOKEN]")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Close the response body
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body
	fmt.Println(string(body))
}
