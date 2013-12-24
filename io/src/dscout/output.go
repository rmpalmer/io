package dscout

import (
	"os"
	"compress/gzip"
	"io"
	"strings"
	"records"
	"formats"
)

type Dscout struct {
	closer		func()
	marshaler	formats.RecordMarshaler
}

func NewDscout (filename string) *Dscout {
	d := new(Dscout)
	file, err := os.Create(filename)
	if err != nil {
		return nil
	}
	d.closer = func() {
		file.Close()
	}
	var writer io.WriteCloser = file
	var compressor *gzip.Writer
	if strings.HasSuffix(filename, ".gz") {
		compressor = gzip.NewWriter(file)
		d.closer = func() { compressor.Close(); file.Close() }
		writer = compressor
	}
	uncompressed_name := strings.TrimRight(filename, ".gz")
	switch {
		case strings.HasSuffix(uncompressed_name, ".gob"):
			d.marshaler = new(formats.GobMarshaler)
		case strings.HasSuffix(uncompressed_name, ".xml"):
			d.marshaler = new(formats.XmlMarshaler)
	}
	if (d.marshaler != nil) {
		d.marshaler.InitFile(writer)
	}
	return d
}

func (d *Dscout) HandleTrace(trace *records.Trace) {
	d.marshaler.MarshalTrace(trace)
}

func (d *Dscout) HandleEod() {
	d.closer()
}

