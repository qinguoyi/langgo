all: demo, registtool

.PHONY: demo
demo:
	GOOS=linux GOARCH=amd64 go build -o bin/demo cmd/demo/main.go

.PHONY: registtool
registtool:
	GOOS=linux GOARCH=amd64 go build -o bin/registtool cmd/registtool/main.go

.PHONY: clean
clean:
	rm -rf bin/