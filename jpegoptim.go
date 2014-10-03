package main

/*
#include "jpegoptim/jpegoptim.h"
extern void jpeg_memory_dest(j_compress_ptr cinfo, unsigned char **bufptr, size_t *bufsizeptr, size_t incsize);
*/
import "C"

func main() {
	C.jpeg_memory_dest()
}
