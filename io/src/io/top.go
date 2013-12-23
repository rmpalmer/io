package main 

import (
	"dscout"
	"io"
	"records"
)

func main() {
	file_1, closer_1, _ := dscout.CreateFile("foo.gz")
//	if closer_1 != nil {
//        defer closer_1()
//    }
	marshaller := dscout.GobMarshaler{}
    var writer io.WriteCloser = file_1
    t := records.NewTrace(5,15)
    for i:=0; i<4; i++ {
    	t.Header[i] = i
    }
    for i:=0; i<10; i++ {
    	t.Data[i] = 10.0 - float64(i) 
    }
    t.Summarize()
    t.Detail()
    marshaller.InitFile(writer)
	marshaller.MarshalTrace(writer,t)
	closer_1()
	
	file_2, closer_2, _ := dscout.OpenFile("foo.gz")
//	if closer_2 != nil {
//        defer closer_2()
//    }
    unmarshaler := dscout.GobMarshaler{}
    var reader io.ReadCloser = file_2
    //unmarshaler.ValidateFile(reader)
	x, _ := unmarshaler.UnmarshalTrace(reader)
	closer_2()
	x.Summarize()
	x.Detail()
}

