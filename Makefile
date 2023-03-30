build:
	go build -o ./bin/pizzeria-management.go ./src/run/pizzeria-management.go

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ./bin/pizzeria-management_linux_amd64 ./src/run/pizzeria-management.go

run:
	go run ./src/run/pizzeria-management.go