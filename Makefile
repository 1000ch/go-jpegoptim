default: build

jpegoptim:
	@mkdir jpegoptim
	@git clone git@github.com:tjko/jpegoptim.git

build: jpegoptim jpegoptim.h jpegoptim.go
	@go tool cgo jpegoptim.go

.PHONY: clean build