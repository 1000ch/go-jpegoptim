package main

/*
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
