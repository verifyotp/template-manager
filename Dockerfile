FROM golang:alpine

WORKDIR /app
COPY . .
RUN go build -o /app/main .

EXPOSE 8080
CMD ["/app/main"]
