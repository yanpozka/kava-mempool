
build:
	go build -o kava main.go

exec:
	./kava

run:
	go run main.go

tests:
	go test -v ./...
