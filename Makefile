APP_NAME = "InstaPurity"

dev:
	go run ./cmd/main.go

run:
	./bin/$(APP_NAME)

build:
	go build -ldflags="-s -w" -gcflags=all=-trimpath=$(PWD) -o ./bin/$(APP_NAME) ./cmd/main.go

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -gcflags=all=-trimpath=$(PWD) -o ./bin/$(APP_NAME).exe ./cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -gcflags=all=-trimpath=$(PWD) -o ./bin/$(APP_NAME).linux ./cmd/main.go
