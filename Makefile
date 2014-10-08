GO=go
REPO=git@github.com:tjko/jpegoptim.git
SRC=jpegoptim-src

default: patch build

$(SRC):
	# prepare source
	@mkdir $(SRC)
	@git clone $(REPO) $(SRC)

build: $(SRC) jpegoptim.go
	# make
	@cd $(SRC);make libjpegoptim.a

	# build
	@$(GO) build jpegoptim.go

patch: $(SRC)
	# configure
	@cd $(SRC);./configure --prefix=`pwd`/local

	# apply patch to jpegoptim/Makefile
	@patch -u $(SRC)/Makefile < Makefile.patch

	# apply patch to jpegoptim/jpegoptim.c
	@patch -u $(SRC)/jpegoptim.c < jpegoptim.c.patch

.PHONY: patch build