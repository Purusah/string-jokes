build:
	go build -o build/app/vowels cmd/vowels/main.go
	go build -o build/app/pig-latin cmd/pig-latin/main.go
	chmod u+x build/app/vowels
	chmod u+x build/app/vowels

test:
	go test -v ./...

.PHONY: build
