
build:
	go build -o bin/server cmd/main.go

run:
	cd bin && ./server

start: build run