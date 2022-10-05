FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM alpine
WORKDIR /app
COPY --from=builder /app/goquestionapi .
COPY app.env .

EXPOSE 8080
CMD [ "/app/goquestionapi" ]
