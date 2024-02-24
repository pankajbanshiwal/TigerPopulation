# Inspired from https://github.com/OT-CONTAINER-KIT/redis-operator/blob/dc07470664a3f504ecc8d95d9194ec01f15fb9fe/Makefile
# NEXT DEV VERSION
DEV_VERSION ?= 0.0.3
# # NEXT STAGE VERSION
STAGE_VERSION ?= 0.0.1
# # NEXT LIVE VERSION
LIVE_VERSION ?= 0.0.1

# # Image URL for different environments
# DEV_IMG ?= 912920125631.dkr.ecr.ap-south-1.amazonaws.com/orion:d$(DEV_VERSION)
# STG_IMG ?= 912920125631.dkr.ecr.ap-south-1.amazonaws.com/orion:s$(STAGE_VERSION)
# LIV_IMG ?= 912920125631.dkr.ecr.ap-south-1.amazonaws.com/orion:p$(LIVE_VERSION)
DEV_IMG ?= pankaj92banshiwal/tigerpopulation2024:d$(DEV_VERSION)
STG_IMG ?= pankaj92banshiwal/tigerpopulation2024:s$(STAGE_VERSION)
LIV_IMG ?= pankaj92banshiwal/tigerpopulation2024:p$(LIVE_VERSION)


run-dev:
	go run -tags dev main.go

run-stg:
	go run -tags stage main.go

run-live:
	go run -tags live main.go

build-dev:
	docker buildx build --platform=linux/arm64 -t ${DEV_IMG} -f dockerfile . && docker push ${DEV_IMG}

build-stg:
	docker buildx build --platform=linux/arm64 -t ${STG_IMG} -f dockerfile.staging . && docker push ${STG_IMG}

build-live:
	docker buildx build --platform=linux/arm64 -t ${LIV_IMG} -f dockerfile.production . && docker push ${LIV_IMG}

push-dev:
	docker push ${DEV_IMG}

push-stg:
	docker push ${STG_IMG}

push-live:
	docker push ${LIV_IMG}
