GO=go
REPO=git@github.com:tjko/jpegoptim.git
SRC=jpegoptim-src

default: patch build

$(SRC):
	@mkdir $(SRC)
	@git clone $(REPO) $(SRC)

build: $(SRC) jpegoptim.go
	# make
	@cd $(SRC);make all

	# build
	@$(GO) build jpegoptim.go

.PHONY: build

patch: $(SRC)
	# configure
	@cd $(SRC);./configure --prefix=`pwd`/local

	# apply patch to jpegoptim/Makefile
	@patch -u $(SRC)/Makefile < Makefile.patch

	# comment out main function in jpegoptim.c
	@sed -e "287,854d" $(SRC)/jpegoptim.c