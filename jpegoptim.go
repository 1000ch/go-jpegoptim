package main

/*
#include <stdio.h>
#include <stdlib.h>
#cgo amd64 386 CFLAGS: -DX86=1
#cgo LDFLAGS: -L/usr/local/Cellar/jpeg/8d/lib/ -L./jpegoptim-src -ljpegoptim -ljpeg
#cgo CFLAGS: -I/usr/local/Cellar/jpeg/8d/include
#include <jpeglib.h>
#include "jpegoptim-src/jpegoptim.h"
extern void jpeg_memory_dest(j_compress_ptr cinfo, unsigned char **bufptr, size_t *bufsizeptr, size_t incsize);
extern int jpegoptim_main(int argc, char **argv);
*/
import "C"
import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println("jpegoptim.go main", args)

	C.jpegoptim_main()
}
