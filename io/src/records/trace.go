package records

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Trace struct {
	Hlen	int
	Dlen	int
	Header  []int
	Data    []float64
}

func NewTrace(h, d int) *Trace {
	t := &Trace{h, d, nil, nil}
	t.Header = make([]int, h)
	t.Data = make([]float64, d)
	return t
}

func (t *Trace) HeaderSize() int {
	return t.Hlen
}

func (t *Trace) DataSize() int {
	return t.Dlen
}


type GobTrace struct {
	Hlen	int
	Dlen	int
	Header  []int
	Data    []float64
}

func (trace *Trace) GobEncode() ([]byte, error) { 
	fmt.Printf("start GobEncode\n")
    gobTrace := GobTrace{trace.Hlen, trace.Dlen,
        trace.Header, trace.Data}
    gobTrace.Summarize()
    var buffer bytes.Buffer
    encoder := gob.NewEncoder(&buffer)
    err := encoder.Encode(gobTrace)
    return buffer.Bytes(), err
}

func (trace *Trace) GobDecode(data []byte) error {
	fmt.Printf("start GobDecode\n")
    var gobTrace GobTrace
    buffer := bytes.NewBuffer(data)
    decoder := gob.NewDecoder(buffer)
    if err := decoder.Decode(&gobTrace); err != nil {
        return err
    }
    gobTrace.Summarize()
    *trace = Trace{gobTrace.Hlen, gobTrace.Dlen,
        gobTrace.Header, gobTrace.Data}
    return nil
}

func (trace *Trace) Summarize() {
	fmt.Printf("trace [%d %d]\n",trace.Hlen, trace.Dlen)
}

func (trace *Trace) Detail() {
	fmt.Printf("trace [")
	for i:=0; i<trace.Hlen; i++ {
		fmt.Printf("%d ",trace.Header[i])
	}
	fmt.Printf("]\n")
	fmt.Printf("      [")
	for i:=0; i<trace.Dlen; i++ {
		fmt.Printf("%f ",trace.Data[i])
	}
	fmt.Printf("]\n")
}

func (trace *GobTrace) Summarize() {
	fmt.Printf("gob trace [%d %d]\n",trace.Hlen, trace.Dlen)
}
	