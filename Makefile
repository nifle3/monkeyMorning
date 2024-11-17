.SILENT:

run:
	go build -o main ./cmd/monkeyMorning/main.go
	./main

lint:
	golangci-lint run ./...

test:
	go build -o main ./cmd/monkeyMorning/main.go
	go test -v ./...