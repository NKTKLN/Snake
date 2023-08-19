build:
	rm -rf bin
	mkdir bin
	GOOS=windows GOARCH=amd64 go build -o bin/snake-amd64-win.exe cmd/snake/main.go 
	GOOS=darwin GOARCH=amd64 go build -o bin/snake-amd64-macos cmd/snake/main.go 
	GOOS=linux GOARCH=amd64 go build -o bin/snake-amd64-linux cmd/snake/main.go 

run:
	go run cmd/snake/main.go 
