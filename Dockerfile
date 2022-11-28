## Multistage build
FROM golang:1.19.2-alpine as build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src
COPY . .
RUN go mod download
RUN go mod tidy
RUN go build -o /api ./cmd/api/main.go

ENTRYPOINT ["/api"]
