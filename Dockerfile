FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o bookish-bliss-api cmd/bookishbliss/main.go

FROM alpine:latest AS migrator

WORKDIR /db/migrations

COPY db/migrations ./

# Download Goose
RUN apk add --no-cache curl && \
    curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh

FROM scratch

WORKDIR /

COPY --from=builder /app/bookish-bliss-api .

CMD ["./bookish-bliss-api"]
