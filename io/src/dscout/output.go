package dscout

import (
	"os"
	"compress/gzip"
	"io"
	"strings"
	"fmt"
)


func CreateFile(filename string) (io.WriteCloser, func(), error) {
	fmt.Printf("opening %s for write\n",filename)
    file, err := os.Create(filename)
    if err != nil {
        return nil, nil, err
    }
    closer := func() {
    	fmt.Printf("closing output file\n") 
    	file.Close() 
    }
    var writer io.WriteCloser = file
    var compressor *gzip.Writer
    if strings.HasSuffix(filename, ".gz") {
        compressor = gzip.NewWriter(file)
        closer = func() { compressor.Close(); file.Close() }
        writer = compressor
    }
    return writer, closer, nil
}
