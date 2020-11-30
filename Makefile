BINARY := cfallow

.PHONY: windows
windows:
	mkdir -p release/windows/
	GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.Version=$(cat .env)' -X 'main.BuildTime=$(date)' -X 'main.Commit=$(git rev-parse --short HEAD)'" -o release/windows/$(BINARY).exe
	tar -cf release/windows.tar.gz release/windows

.PHONY: linux
linux:
	mkdir -p release/linux/
	GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=$(cat .env)' -X 'main.BuildTime=$(date)' -X 'main.Commit=$(git rev-parse --short HEAD)'" -o release/linux/$(BINARY)
	tar -cf release/linux.tar.gz release/linux

.PHONY: darwin
darwin:
	mkdir -p release/darwin/
	GOOS=darwin GOARCH=amd64 go build -ldflags="-X 'main.Version=$(cat .env)' -X 'main.BuildTime=$(date)' -X 'main.Commit=$(git rev-parse --short HEAD)'" -o release/darwin/$(BINARY)
	tar -cf release/darwin.tar.gz release/darwin

.PHONY: build
build:  windows linux darwin

.PHONY: sign
sign:
	# sign
	gon -log-level=debug -log-json ./gon.json

.PHONY: release
release:  sign
