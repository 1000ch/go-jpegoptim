CGO=go tool cgo
SRC=jpegoptim-src

default: build

$(SRC):
	@mkdir $(SRC)
	@git clone git@github.com:tjko/jpegoptim.git $(SRC)

build: $(SRC) jpegoptim.h jpegoptim.go
	@$(CGO) -debug-gcc jpegoptim.go

.PHONY: clean build