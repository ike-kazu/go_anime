FROM golang:1.18.3-alpine

WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go mod tidy

EXPOSE 8080