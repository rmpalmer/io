package dscout

import (
	"io"
	"encoding/gob"
	"records"
	"errors"
	"fmt"
)

type GobMarshaler struct{}

func (GobMarshaler) InitFile(writer io.Writer) error {
    encoder := gob.NewEncoder(writer)
    if err := encoder.Encode(magicNumber); err != nil {
        return err
    }
    if err := encoder.Encode(fileVersion); err != nil {
        return err
    }
    return nil
}

func (GobMarshaler) ValidateFile(reader io.Reader) (error) {
    decoder := gob.NewDecoder(reader)
    var magic int
    if err := decoder.Decode(&magic); err != nil {
        return err
    }
    if magic != magicNumber {
        return errors.New("cannot read non-trace gob file")
    } else {
    	fmt.Printf("read magic number %d\n",magic)
    }
    var version int
    if err := decoder.Decode(&version); err != nil {
        return err
    }
    if version > fileVersion {
        return fmt.Errorf("version %d is too new to read", version)
    } else {
    	fmt.Printf("read file version %d\n",version)
    }
    fmt.Printf("ValidateFile no errors\n")
	return nil
}

func (GobMarshaler) MarshalTrace(writer io.Writer,
    trace *records.Trace) error {
    fmt.Printf("starting MarshalTrace\n")
    encoder := gob.NewEncoder(writer)
//    if err := encoder.Encode(magicNumber); err != nil {
//        return err
//    }
//    if err := encoder.Encode(fileVersion); err != nil {
//        return err
//    }
//    fmt.Printf("marshaller about to encode trace\n")
    err := encoder.Encode(trace)
    fmt.Printf("done calling encoder.Encode %s\n",err)
    return err
}

func (GobMarshaler) UnmarshalTrace(reader io.Reader) (*records.Trace,
    error) {
    decoder := gob.NewDecoder(reader)
    fmt.Printf("starting UnmarshalTrace\n")
    var magic int
    if err := decoder.Decode(&magic); err != nil {
        return nil, err
    }
    if magic != magicNumber {
        return nil, errors.New("cannot read non-trace gob file")
    } else {
    	fmt.Printf("read magic number %d\n",magic)
    }
    var version int
    if err := decoder.Decode(&version); err != nil {
        return nil, err
    }
    if version > fileVersion {
        return nil, fmt.Errorf("version %d is too new to read", version)
    } else {
    	fmt.Printf("read file version %d\n",version)
    }
    var trace records.Trace
    fmt.Printf("unmarshaller about to decode trace\n")
    err := decoder.Decode(&trace)
    fmt.Printf("done calling decoder.Decode %s\n",err)
    return &trace, err
}
