all: langlangGo

.PHONY: langlangGo
langlangGo:
	GOOS=linux GOARCH=amd64 go build -o bin/langlangGo cmd/demo/main.go

.PHONY: clean
clean:
	rm -rf bin/