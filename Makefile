VERSION     := 1.0.0
LDFLAGS     := -ldflags="-s -w -X main.version=$(VERSION)"
BUILD_FLAGS := CGO_ENABLED=0

.PHONY: all linux windows clean

all: linux windows

linux:
	$(BUILD_FLAGS) GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/duofetch-linux-amd64 .
	$(BUILD_FLAGS) GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/duofetch-linux-arm64 .

windows:
	$(BUILD_FLAGS) GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/duofetch-windows-amd64.exe .

# Build for the current platform (default for development).
build:
	$(BUILD_FLAGS) go build $(LDFLAGS) -o duofetch .

clean:
	rm -rf dist/ duofetch duofetch.exe
