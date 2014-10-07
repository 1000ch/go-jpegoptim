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
import (
	"fmt"
	"log"
	"errors"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println("jpegoptim.go main", args)

	if !os.IsExist(args[0]) {
		log.Fatal("argument file does not exist")
	}

	file, err := os.Open(args[0])
	if err != nil {
		log.fatal(err)
	}
	defer file.Close()


	// cinfo := C.struct_j_compress_ptr{}
	// C.jpeg_memory_dest(&cinfo)
}
