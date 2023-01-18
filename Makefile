version?=0.0.0
VERSION?=$(version)
export VERSION

test:
	go test ./...

clean:
	rm -rf ./bin

build:
	$(MAKE) build-darwin
	$(MAKE) build-linux
	$(MAKE) build-windows

build-darwin:
	GOOS=darwin GOARCH=amd64 \
		go build -o ./bin/subenv-$(VERSION)-darwin-amd64/subenv \
		-ldflags "-X 'main.version=$(VERSION)' -X 'main.goos=darwin' -X 'main.goarch=amd64'" \
		./cmd/cli/subenv.go \
		&& cd bin; zip -r subenv-$(VERSION)-darwin-amd64.zip ./subenv-$(VERSION)-darwin-amd64; cd ..

build-linux:
	GOOS=linux GOARCH=amd64 \
		go build -o ./bin/subenv-$(VERSION)-linux-amd64/subenv \
		-ldflags "-X 'main.version=$(VERSION)' -X 'main.goos=linux' -X 'main.goarch=amd64'" \
		./cmd/cli/subenv.go \
		&& cd bin; tar -czf subenv-$(VERSION)-linux-amd64.tar.gz ./subenv-$(VERSION)-linux-amd64; cd ..

build-windows:
	GOOS=windows GOARCH=amd64 \
		go build -o ./bin/subenv-$(VERSION)-windows-amd64/subenv.exe \
		-ldflags "-X 'main.version=$(VERSION)' -X 'main.goos=windows' -X 'main.goarch=amd64'" \
		./cmd/cli/subenv.go \
		&& cd bin; zip -r subenv-$(VERSION)-windows-amd64.zip ./subenv-$(VERSION)-windows-amd64; cd ..

all: test clean build
