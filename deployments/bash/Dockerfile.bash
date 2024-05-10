# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bashExecAPI ./cmd/server.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

USER root

WORKDIR /

COPY --from=build-stage /bashExecAPI /bashExecAPI

EXPOSE 8000

# need to mount volume into /configs/config.yaml
ENTRYPOINT ["/bashExecAPI"]