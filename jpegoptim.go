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
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Println("jpegoptim.go main", args)

	if len(args) == 0 {
		log.Fatal("There is no argument")
	}

	_, err := os.Stat(args[0])
	if os.IsNotExist(err) {
		log.Fatal("Specified file does not exist")
	}

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//var cinfo C.struct_jpeg_compress_struct
	//var bufptr *C.uchar
	//var bufsizeptr C.size_t
	//var incsize C.size_t

	//C.jpeg_memory_dest(&cinfo, &bufptr, &bufsizeptr, &incsize)
}
