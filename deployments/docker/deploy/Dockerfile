FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/app cmd/fiesta/fiesta.go

EXPOSE 8080 8080

CMD ["./bin/app"]
