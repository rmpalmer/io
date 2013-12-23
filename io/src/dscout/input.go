package dscout

import (
	"os"
	"io"
	"compress/gzip"
	"strings"
	"fmt"
)

func OpenFile(filename string) (io.ReadCloser, func(), error) {
	fmt.Printf("opening %s for read\n",filename)
    file, err := os.Open(filename)
    if err != nil {
        return nil, nil, err
    }
    closer := func() {
    	fmt.Printf("closing input file\n")
    	file.Close() 
   	}
    var reader io.ReadCloser = file
    var decompressor *gzip.Reader
    if strings.HasSuffix(filename, ".gz") {
        if decompressor, err = gzip.NewReader(file); err != nil {
            return file, closer, err
        }
        closer = func() { decompressor.Close(); file.Close() }
        reader = decompressor
    }
    return reader, closer, nil
}
