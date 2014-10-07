GO=go
CGO=go tool cgo
REPO=git@github.com:tjko/jpegoptim.git
SRC=jpegoptim-src

default: build

$(SRC):
	@mkdir $(SRC)
	@git clone $(REPO) $(SRC)

build: $(SRC) jpegoptim.h jpegoptim.go
	# configure
	@cd $(SRC);./configure --prefix=`pwd`/local

	# apply patch
	@patch -u $(SRC)/Makefile < Makefile.patch

	# comment out main function in jpegoptim.c
	@sed -e "287,854d" $(SRC)/jpegoptim.c

	# make
	@cd $(SRC);make all

	# build
	@$(GO) build jpegoptim.go

.PHONY: clean build