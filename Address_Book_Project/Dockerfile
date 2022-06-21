FROM golang:1.18
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8081
CMD ["go", "run", "server/main.go"]
