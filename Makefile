BuildCmd = CGO_ENABLED=0 GOARCH=amd64 go build -o _ ./tests/build.go

.PHONY: build
build: build-sys build-win build-linux build-darwin

# build program with current system
.PHONY: build-sys
build-sys:
	go build -o bin-sys ./tests/build.go

.PHONY: build-win
build-win:
	GOOS=windows $(patsubst _, bin-win.exe, $(BuildCmd))

.PHONY: build-linux
build-linux:
	GOOS=windows $(patsubst _, bin-linux, $(BuildCmd))

.PHONY: build-darwin
build-darwin:
	GOOS=windows $(patsubst _, bin-darwin, $(BuildCmd))

.PHONY: clean
clean:
	rm -f ./bin-*