FROM golang@latest
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags all="-w -s" -o . ./...
CMD ["./client"]