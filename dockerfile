# First Stage
# https://www.docker.com/blog/faster-multi-platform-builds-dockerfile-cross-compilation-guide/

FROM --platform=$BUILDPLATFORM golang:1.21-alpine
# Magic line, notice in use that the lib name is different!
RUN apk update && apk add ca-certificates git

RUN mkdir /App
ADD . /App/
WORKDIR /App
ARG TARGETOS TARGETARCH
# https://github.com/golang/go/wiki/GoArm
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -tags dev -o server .

# Second Stage
FROM --platform=$BUILDPLATFORM alpine:latest
WORKDIR /root/
RUN apk add --no-cache tzdata
CMD ["/App"]
# Copy from first stage
COPY --from=0 /App/server /App
COPY --from=0 /etc/ssl/certs /etc/ssl/certs
COPY config.yaml .
