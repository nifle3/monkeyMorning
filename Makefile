.SILENT:

build:
	go build -o build/mainMorning ./cmd/monkeyMorning/main.go
	go build -o build/mainAfternoon ./cmd/monkeyAfternoon/main.go
	go build -o build/mainEvening ./cmd/monkeyEvening/main.go

runmorning:
	go build -o main ./cmd/monkeyMorning/main.go
	./main

runafternoon:
	go build -o main ./cmd/monkeyAfternoon/main.go
	./main

runevening:
	go build -o main ./cmd/monkeyEvening/main.go
	./main

generate:
	sqlc vet
	sqlc verify
	sqlc generate

lint:
	golangci-lint run ./...

test: build generate lint
	go test -v ./...

deps:
	go mod tidy
	go mod vendor

setup:
