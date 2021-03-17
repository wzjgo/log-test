run:
	go run main.go
build-cross:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go	