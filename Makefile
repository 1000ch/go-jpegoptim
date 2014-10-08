GO=go
REPO=git@github.com:tjko/jpegoptim.git
SRC=jpegoptim-src

default: build

$(SRC):
	@mkdir $(SRC)
	@git clone $(REPO) $(SRC)

patch:
	# configure
	@cd $(SRC);./configure --prefix=`pwd`/local

	# apply patch to jpegoptim/Makefile
	@patch -u $(SRC)/Makefile < Makefile.patch

	# comment out main function in jpegoptim.c
	@sed -e "287,854d" $(SRC)/jpegoptim.c

build: $(SRC) jpegoptim.go
	# make
	@cd $(SRC);make all

	# build
	@$(GO) build jpegoptim.go

.PHONY: build