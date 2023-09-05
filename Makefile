linux:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o ./bin/evalgpt ./*.go

linux-arm:
	mkdir -p bin
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -mod=vendor -o ./bin/evalgpt ./*.go

darwin-arm:
	mkdir -p bin
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -mod=vendor -o ./bin/evalgpt ./*.go

darwin:
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o ./bin/evalgpt ./*.go

clean:
	rm -rf bin

fmt:
	go fmt ./...