#ifndef _GO_JPEGOPTIM_H
#define _GO_JPEGOPTIM_H

#ifdef __cplusplus
extern "C" {
#endif

    void jpeg_memory_dest (j_compress_ptr cinfo, unsigned char **bufptr, size_t *bufsizeptr, size_t incsize);

#ifdef __cplusplus
}
#endif

#endif