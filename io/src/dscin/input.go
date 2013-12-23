package dscin

import (
	"os"
	"io"
	"compress/gzip"
	"strings"
	"fmt"
	"records"
	"formats"
)

type Dscin struct {
	closer		func()
	marshaler	formats.GobMarshaler
}

func NewDscin(filename string) *Dscin {
	d := new(Dscin)
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	d.closer = func() {
		fmt.Printf("closing file %s\n", filename)
		file.Close()
	}
	var reader io.ReadCloser = file
	var decompressor *gzip.Reader
	if strings.HasSuffix(filename, ".gz") {
		decompressor, _ = gzip.NewReader(file)
		d.closer = func() { decompressor.Close(); file.Close() }
		reader = decompressor
	}
	d.marshaler = formats.GobMarshaler{}
	d.marshaler.ValidateFile(reader)
	return d
}

func (d *Dscin) HandleTrace() *records.Trace {
	t, err := d.marshaler.UnmarshalTrace()
	if (err != nil) { 
		return nil
	} else {
		return t
	}
}
