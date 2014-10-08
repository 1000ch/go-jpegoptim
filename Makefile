GO=go
REPO=git@github.com:tjko/jpegoptim.git
SRC=jpegoptim-src

default: patch build

$(SRC):
	@echo "+++source preparing+++"
	@mkdir $(SRC)
	@git clone $(REPO) $(SRC)

build: $(SRC) jpegoptim.go
	@echo "+++building+++"
	# make
	@cd $(SRC);make all

	# build
	@$(GO) build jpegoptim.go

patch: $(SRC)
	@echo "+++patch applying+++"

	# configure
	@cd $(SRC);./configure --prefix=`pwd`/local

	# apply patch to jpegoptim/Makefile
	@patch -u $(SRC)/Makefile < Makefile.patch

	# apply patch to jpegoptim/jpegoptim.c
	@patch -u $(SRC)/jpegoptim.c < jpegoptim.c.patch

.PHONY: patch build