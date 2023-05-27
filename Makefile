all:
	go mod tidy
	go fmt ./...
	go test -cover ./...
	go install ./...

test:
	go mod tidy
	go test -v --coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
