build:
	go build -o ./bin/alloy.exe main.go
run: build
	./bin/alloy.exe