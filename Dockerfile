FROM golang:latest

WORKDIR /opt/api

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

EXPOSE 8080

ENTRYPOINT ["go", "run", "main.go"]
