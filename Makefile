default: build

jpegoptim-src:
	@mkdir jpegoptim-src
	@git clone git@github.com:tjko/jpegoptim.git jpegoptim-src

build: jpegoptim-src jpegoptim.h jpegoptim.go
	@go tool cgo jpegoptim.go

.PHONY: clean build