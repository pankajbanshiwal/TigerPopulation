# To deploy this service on cloud , one has to follow below steps -> Post docker image is pushed to cloud registry , deploy the image to any of the cloud service provider like AWS,GCP etc.

# Selecting an env while running
https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags

https://medium.com/@tharun208/build-tags-in-go-f21ccf44a1b8
## dev
`$ go run -tags dev main.go`
## stage
`go run -tags stage main.go`
## Live
`go run -tags live main.go`


# selecting a env while building
Select appropriate version before building
https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
## dev
`$ docker buildx build --platform=linux/arm64 -t pankaj92banshiwal/tigerpopulation2024:d0.0.0 -f dockerfile .`
## stage
`$ docker buildx build --platform=linux/arm64 -t pankaj92banshiwal/tigerpopulation2024:s0.0.0 -f dockerfile.staging .`
## live
`$ docker buildx build --platform=linux/arm64 -t pankaj92banshiwal/tigerpopulation2024:p0.0.0 -f dockerfile.production .`


# Uploading
`$ docker push pankaj92banshiwal/tigerpopulation2024:<version>`


# docker commands for testing
for run in locally after build the image =>
`docker run -it -p 1008:1008 docker push pankaj92banshiwal/tigerpopulation2024:d0.0.1`  // ideal one

# consuming make file
## running service in dev env
`make run-dev`
## running service in stage env
`make run-stg`
## running service in live
`make run-live`

## building dev docker image and pushing to docker hub
`make build-dev`
## building stage docker image and pushing to docker hub
`make build-stg`
## building live docker image and pushing to docker hub
`make build-live`

for stop the docker image =>
`docker ps`  // `docker stop < containerid >`

# Limitations
`$ Because of short deadline i have stored username password in plain text , we can use pgcrypto module to encrypt and store them safely eg :-   crypt('password', gen_salt('salt_key'))`

`$ I have assumed sightings images to be uploaded by UI to cloud storage and image urls to be stored in back end`

# APIs postman collection for testing -> import the collection to be able to test the APIs
`https://api.postman.com/collections/8163204-1c677a80-8783-444c-8d17-41034af4b538?access_key=PMAT-01HQDRNXNV1D67YYXVXKWFC5DB`

# API Documention for UI developers to be consumed
`https://documenter.getpostman.com/view/8163204/2sA2rCUhKu`


### Who do I talk to? ###
* Pankaj Banshiwal: pankajofcbanshiwal@gmail.com
