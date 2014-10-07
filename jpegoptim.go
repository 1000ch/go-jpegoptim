package main

/*
#include <stdio.h>
#cgo amd64 386 CFLAGS: -DX86=1
#cgo LDFLAGS: -L/usr/local/Cellar/jpeg/8d/lib/ -L./jpegoptim-src -ljpegoptim -ljpeg
#cgo CFLAGS: -I/usr/local/Cellar/jpeg/8d/include
#include <jpeglib.h>
#include "jpegoptim-src/jpegoptim.h"
extern void jpeg_memory_dest(j_compress_ptr cinfo, unsigned char **bufptr, size_t *bufsizeptr, size_t incsize);
*/
import "C"
import "fmt"

func main() {
	fmt.Println("jpegoptim.go main")
}
