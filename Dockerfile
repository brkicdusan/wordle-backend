# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /main cmd/api/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN make test

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main
COPY --from=build-stage /app/assets /assets

USER nonroot:nonroot

EXPOSE 9001

ENTRYPOINT ["/main"]

