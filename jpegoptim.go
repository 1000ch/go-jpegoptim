package main

/*
#cgo amd64 386 CFLAGS: -DX86=1
#cgo LDFLAGS: -L/usr/local/Cellar/jpeg-turbo/1.3.1/lib
#cgo CPPFLAGS: -I/usr/local/Cellar/jpeg-turbo/1.3.1/include
#include <stdio.h>
#include "jpegoptim-src/jpegoptim.h"
extern void jpeg_memory_dest(j_compress_ptr cinfo, unsigned char **bufptr, size_t *bufsizeptr, size_t incsize);
*/
import "C"
import "fmt"

func main() {
	C.jpeg_memory_dest()
	fmt.Println("jpegoptim.go main")
}
