FROM golang:latest as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/api cmd/main.go

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch

WORKDIR /app
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/bin/api .
COPY --from=builder /app/.env .

ENTRYPOINT [ "./api" ]
