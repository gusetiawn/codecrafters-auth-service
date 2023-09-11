FROM golang:1.21.0

WORKDIR /codecrafters-auth-service
COPY . .

RUN go build -o /app/my-app

EXPOSE 8080

CMD ["/app/my-app", "serve"]
